package projectinitialize

import (
	"context"
	"fmt"

	projectset "github.com/openshift/client-go/project/clientset/versioned/typed/project/v1"
	redhatcopv1alpha1 "github.com/redhat-cop/project-initialize-operator/pkg/apis/redhatcop/v1alpha1"
	"github.com/redhat-cop/project-initialize-operator/pkg/controller/logging"
	projectinit "github.com/redhat-cop/project-initialize-operator/pkg/controller/projectinitialize/ocp/project"
	project "github.com/redhat-cop/project-initialize-operator/pkg/controller/projectinitialize/resources"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_projectinitialize")

/**
* USER ACTION REQUIRED: This is a scaffold file intended for the user to modify with their own Controller
* business logic.  Delete these comments after modifying this file.*
 */

// Add creates a new ProjectInitialize Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	client, err := projectset.NewForConfig(mgr.GetConfig())
	if err != nil {
		return nil
	}
	return &ReconcileProjectInitialize{client: mgr.GetClient(), scheme: mgr.GetScheme(), projectClient: client}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("projectinitialize-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ProjectInitialize
	err = c.Watch(&source.Kind{Type: &redhatcopv1alpha1.ProjectInitialize{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// TODO(user): Modify this to be the types you create that are owned by the primary resource
	// Watch for changes to secondary resource Pods and requeue the owner ProjectInitialize
	err = c.Watch(&source.Kind{Type: &corev1.Pod{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &redhatcopv1alpha1.ProjectInitialize{},
	})
	if err != nil {
		return err
	}

	return nil
}

// blank assignment to verify that ReconcileProjectInitialize implements reconcile.Reconciler
var _ reconcile.Reconciler = &ReconcileProjectInitialize{}

// ReconcileProjectInitialize reconciles a ProjectInitialize object
type ReconcileProjectInitialize struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client        client.Client
	scheme        *runtime.Scheme
	projectClient *projectset.ProjectV1Client
}

// Reconcile reads that state of the cluster for a ProjectInitialize object and makes changes based on the state read
// and what is in the ProjectInitialize.Spec
// TODO(user): Modify this Reconcile function to implement your Controller logic.  This example creates
// a Pod as an example
// Note:
// The Controller will requeue the Request to be processed again if the returned error is non-nil or
// Result.Requeue is true, otherwise upon completion it will remove the work from the queue.
func (r *ReconcileProjectInitialize) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ProjectInitialize")
	// Fetch the ProjectInitialize instance
	instance := &redhatcopv1alpha1.ProjectInitialize{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	/* TODO - Add reconcile cycle */
	//Does the project exist?
	projectName := projectinit.GetProjectName(instance.Spec.Team, instance.Spec.Env)
	found, err := r.projectClient.Projects().Get(projectName, metav1.GetOptions{})
	// If project doesn't exist, create it
	if err != nil {
		projectRequest := project.GetProjectRequest(projectName, instance.Spec.DisplayName, instance.Spec.Desc)
		newProject, err := projectinit.InitializeProjectOCP(r.projectClient, projectRequest)
		if err != nil {
			return reconcile.Result{}, err
		}

		// TODO setup ArgoCD, Qoutas, GIT and LDAP
		if instance.Spec.QuotaSize != "" {
			err, quotaSize := projectinit.GetQuotaSizeFromCluster(r.client, instance.Spec.QuotaSize)
			if err != nil {
				return reconcile.Result{}, err
				// TODO reverse the project creation?
			}
			quota := project.GetQuotaResource(fmt.Sprintf("%s-quota", instance.Spec.Team), projectName, *quotaSize)
			err = projectinit.AddQuotaToProject(r.client, quota)
			if err != nil {
				return reconcile.Result{}, err
				// TODO reverse the project creation?
			}
		}
		logging.Log.Info(fmt.Sprintf("Created new project %s", newProject.Name))
	} else {
		logging.Log.Info(fmt.Sprintf("Found project %s", found.Name))
		// Check if labels or annotations have changed
		if instance.Spec.NamespaceDetails != nil {
			err = projectinit.UpdateNamespaceAnnotations(r.client, projectName, instance.Spec.NamespaceDetails)
			if err != nil {
				return reconcile.Result{}, err
			}
		}
	}

	return reconcile.Result{}, nil
}
