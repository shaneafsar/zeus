package do_nvme

const (
	DoStorageClass = "nvme-ssd-block"
	DoNvmePath     = "/dev/vda1"
)

func AddDoNvmeLabels(labels map[string]string) map[string]string {
	labels["fast-disk-node"] = "pv-raid"
	return labels
}
