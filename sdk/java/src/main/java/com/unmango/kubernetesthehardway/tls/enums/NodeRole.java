// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tls.enums;

import com.pulumi.core.annotations.EnumType;
import java.lang.String;
import java.util.Objects;
import java.util.StringJoiner;

    /**
     * The role a node will play in the final cluster
     * 
     */
    @EnumType
    public enum NodeRole {
        Controlplane("controlplane"),
        Worker("worker");

        private final String value;

        NodeRole(String value) {
            this.value = Objects.requireNonNull(value);
        }

        @EnumType.Converter
        public String getValue() {
            return this.value;
        }

        @Override
        public String toString() {
            return new StringJoiner(", ", "NodeRole[", "]")
                .add("value='" + this.value + "'")
                .toString();
        }
    }
