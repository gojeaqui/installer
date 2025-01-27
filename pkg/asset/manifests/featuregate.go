package manifests

import (
	"path/filepath"

	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/yaml"

	configv1 "github.com/openshift/api/config/v1"
	"github.com/openshift/installer/pkg/asset"
	"github.com/openshift/installer/pkg/asset/installconfig"
)

var fgFileName = filepath.Join(openshiftManifestDir, "99_feature-gate.yaml")

// FeatureGate generates the feature gate manifest.
type FeatureGate struct {
	FileList []*asset.File
	Config   configv1.FeatureGate
}

var _ asset.WritableAsset = (*Proxy)(nil)

// Name returns a human-friendly name for the asset.
func (*FeatureGate) Name() string {
	return "Feature Gate Config"
}

// Dependencies returns all of the dependencies directly needed to generate
// the asset.
func (*FeatureGate) Dependencies() []asset.Asset {
	return []asset.Asset{
		&installconfig.InstallConfig{},
	}
}

// Generate generates the FeatureGate CRD.
func (f *FeatureGate) Generate(dependencies asset.Parents) error {
	installConfig := &installconfig.InstallConfig{}
	dependencies.Get(installConfig)

	// A FeatureGate could be populated on every install,
	// even for those using the default feature set, but for
	// continuity let's only create a cluster feature gate
	// when non-default feature gates are enabled.
	if installConfig.Config.FeatureSet != configv1.Default {
		f.Config = configv1.FeatureGate{
			TypeMeta: metav1.TypeMeta{
				APIVersion: configv1.SchemeGroupVersion.String(),
				Kind:       "FeatureGate",
			},
			ObjectMeta: metav1.ObjectMeta{
				Name: "cluster",
			},
			Spec: configv1.FeatureGateSpec{
				FeatureGateSelection: configv1.FeatureGateSelection{
					FeatureSet: installConfig.Config.FeatureSet,
				},
			},
		}

		configData, err := yaml.Marshal(f.Config)
		if err != nil {
			return errors.Wrapf(err, "failed to create %s manifests from InstallConfig", f.Name())
		}

		f.FileList = []*asset.File{
			{
				Filename: fgFileName,
				Data:     configData,
			},
		}
	}
	return nil
}

// Files returns the files generated by the asset.
func (f *FeatureGate) Files() []*asset.File {
	return f.FileList
}

// Load loads the already-rendered files back from disk.
func (f *FeatureGate) Load(ff asset.FileFetcher) (bool, error) {
	return false, nil
}
