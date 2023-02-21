#include <stdio.h>
#include <string.h>
#include "cgroup.h"

int get_cgroup_data_by_cgo(const char *pid_cgroup, char *pod_uid, char *container_id, size_t size) {
    struct get_cgroup_data_with_containerd_return ret = get_cgroup_data_with_containerd(pid_cgroup);
    if (ret.r0 == NULL || ret.r1 == NULL) {
        return 1;
    }
    strncpy(pod_uid, ret.r0, size);
    pod_uid[size - 1] = '\0';
    strncpy(container_id, ret.r1, size);
    pod_uid[size - 1] = '\0';
    return 0;
}

int main() {
    char pid_cgroup[] = "./tests/containerd-systemd";
    char pod_uid[4096], container_id[4096];
    if (get_cgroup_data_by_cgo(pid_cgroup, pod_uid, container_id, sizeof(container_id))) {
        printf("faild get cgroup data by cgo\n");
        return 1;
    }
    printf("pod_uid: %s\n", pod_uid);
    printf("container_id: %s\n", container_id);
    return 0;
}