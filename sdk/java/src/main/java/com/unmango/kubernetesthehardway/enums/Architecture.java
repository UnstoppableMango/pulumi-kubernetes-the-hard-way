// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.enums;

import com.pulumi.core.annotations.EnumType;
import java.lang.String;
import java.util.Objects;
import java.util.StringJoiner;

    /**
     * CPU architecture
     * 
     */
    @EnumType
    public enum Architecture {
        Amd64("amd64"),
        Arm64("arm64");

        private final String value;

        Architecture(String value) {
            this.value = Objects.requireNonNull(value);
        }

        @EnumType.Converter
        public String getValue() {
            return this.value;
        }

        @Override
        public String toString() {
            return new StringJoiner(", ", "Architecture[", "]")
                .add("value='" + this.value + "'")
                .toString();
        }
    }
