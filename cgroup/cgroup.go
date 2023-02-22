package main

import (
	"os"
	"strings"
)

import "C"

func main() {
}

func normalize(id string) string {
	return strings.ReplaceAll(id, "_", "-")
}

//export get_cgroup_data_with_containerd
func get_cgroup_data_with_containerd(path *C.char) (pod_id *C.char, container_id *C.char) {
	d, err := os.ReadFile(C.GoString(path))
	if err != nil {
		return nil, nil
	}
	lines := strings.Split(string(d), "\n")
	var rawLine string
	// find line name=systemd
	// /proc/self/cgroup
	// 1:name=systemd:/system.slice/containerd.service/kubepods-pod3afbda42_dabf_482d_962e_77bada079c54.slice:cri-containerd:68ac51f452910c79d29c2f16d5130432d30e94b890069195d8b2381b88e11489
	for _, line := range lines {
		if strings.Contains(line, "name=systemd") {
			rawLine = line
			break
		}
	}
	// find pod ID and container ID
	raws := strings.SplitN(rawLine, "-pod", 2)
	if len(raws) == 2 {
		ids := strings.SplitN(raws[1], ".slice:cri-containerd:", 2)
		if len(ids) == 2 {
			podID := normalize(ids[0])
			containerID := ids[1]
			return C.CString(podID), C.CString(containerID)
		}
	}
	return nil, nil
}
