// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tls.enums;

import com.pulumi.core.annotations.EnumType;
import java.lang.String;
import java.util.Objects;
import java.util.StringJoiner;

    /**
     * Private key algorithm.
     * 
     */
    @EnumType
    public enum Algorithm {
        RSA("RSA"),
        ECDSA("ECDSA"),
        ED25519("ED25519");

        private final String value;

        Algorithm(String value) {
            this.value = Objects.requireNonNull(value);
        }

        @EnumType.Converter
        public String getValue() {
            return this.value;
        }

        @Override
        public String toString() {
            return new StringJoiner(", ", "Algorithm[", "]")
                .add("value='" + this.value + "'")
                .toString();
        }
    }
