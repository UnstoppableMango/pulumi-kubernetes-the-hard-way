// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools.enums;

import com.pulumi.core.annotations.EnumType;
import java.lang.String;
import java.util.Objects;
import java.util.StringJoiner;

    @EnumType
    public enum HostnamectlJsonMode {
        Short_("short"),
        Pretty("pretty"),
        Off("off");

        private final String value;

        HostnamectlJsonMode(String value) {
            this.value = Objects.requireNonNull(value);
        }

        @EnumType.Converter
        public String getValue() {
            return this.value;
        }

        @Override
        public String toString() {
            return new StringJoiner(", ", "HostnamectlJsonMode[", "]")
                .add("value='" + this.value + "'")
                .toString();
        }
    }
