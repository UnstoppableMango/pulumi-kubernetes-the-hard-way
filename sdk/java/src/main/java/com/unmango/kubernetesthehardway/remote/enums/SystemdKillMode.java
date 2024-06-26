// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote.enums;

import com.pulumi.core.annotations.EnumType;
import java.lang.String;
import java.util.Objects;
import java.util.StringJoiner;

    /**
     * https://www.freedesktop.org/software/systemd/man/latest/systemd.kill.html#Description
     * 
     */
    @EnumType
    public enum SystemdKillMode {
        Controlgroup("control-group"),
        Mixed("mixed"),
        Process("process"),
        None("none");

        private final String value;

        SystemdKillMode(String value) {
            this.value = Objects.requireNonNull(value);
        }

        @EnumType.Converter
        public String getValue() {
            return this.value;
        }

        @Override
        public String toString() {
            return new StringJoiner(", ", "SystemdKillMode[", "]")
                .add("value='" + this.value + "'")
                .toString();
        }
    }
