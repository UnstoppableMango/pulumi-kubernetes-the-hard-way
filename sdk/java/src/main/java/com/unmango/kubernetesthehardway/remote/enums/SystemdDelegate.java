// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote.enums;

import com.pulumi.core.annotations.EnumType;
import java.lang.String;
import java.util.Objects;
import java.util.StringJoiner;

    /**
     * https://www.man7.org/linux/man-pages/man5/systemd.resource-control.5.html
     * 
     */
    @EnumType
    public enum SystemdDelegate {
        Yes("yes"),
        No("no"),
        Cpu("cpu"),
        Cpuacct("cpuacct"),
        Cpuset("cpuset"),
        Io("io"),
        Blkio("blkio"),
        Memory("memory"),
        Devices("devices"),
        Pids("pids"),
        Bpffirewall("bpf-firewall"),
        Bpfdevices("bpf-devices");

        private final String value;

        SystemdDelegate(String value) {
            this.value = Objects.requireNonNull(value);
        }

        @EnumType.Converter
        public String getValue() {
            return this.value;
        }

        @Override
        public String toString() {
            return new StringJoiner(", ", "SystemdDelegate[", "]")
                .add("value='" + this.value + "'")
                .toString();
        }
    }
