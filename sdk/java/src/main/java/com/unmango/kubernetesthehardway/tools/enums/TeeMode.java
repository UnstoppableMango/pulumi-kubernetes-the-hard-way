// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools.enums;

import com.pulumi.core.annotations.EnumType;
import java.lang.String;
import java.util.Objects;
import java.util.StringJoiner;

    @EnumType
    public enum TeeMode {
        Warn("warn"),
        Warnnopipe("warn-nopipe"),
        Exit("exit"),
        Exitnopipe("exit-nopipe");

        private final String value;

        TeeMode(String value) {
            this.value = Objects.requireNonNull(value);
        }

        @EnumType.Converter
        public String getValue() {
            return this.value;
        }

        @Override
        public String toString() {
            return new StringJoiner(", ", "TeeMode[", "]")
                .add("value='" + this.value + "'")
                .toString();
        }
    }
