// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.remote.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SystemdUnitSection {
    /**
     * @return Configures requirement dependencies, very similar in style to Requires=.
     * 
     */
    private @Nullable List<String> bindsTo;
    /**
     * @return A short human readable title of the unit.
     * 
     */
    private @Nullable String description;
    /**
     * @return A space-separated list of URIs referencing documentation for this unit or its configuration.
     * 
     */
    private @Nullable List<String> documentation;
    /**
     * @return Similar to Wants=, but declares a stronger requirement dependency.
     * 
     */
    private @Nullable List<String> requires;
    /**
     * @return Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
     * 
     */
    private @Nullable List<String> requisite;
    /**
     * @return Configures (weak) requirement dependencies on other units.
     * 
     */
    private @Nullable List<String> wants;

    private SystemdUnitSection() {}
    /**
     * @return Configures requirement dependencies, very similar in style to Requires=.
     * 
     */
    public List<String> bindsTo() {
        return this.bindsTo == null ? List.of() : this.bindsTo;
    }
    /**
     * @return A short human readable title of the unit.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return A space-separated list of URIs referencing documentation for this unit or its configuration.
     * 
     */
    public List<String> documentation() {
        return this.documentation == null ? List.of() : this.documentation;
    }
    /**
     * @return Similar to Wants=, but declares a stronger requirement dependency.
     * 
     */
    public List<String> requires() {
        return this.requires == null ? List.of() : this.requires;
    }
    /**
     * @return Similar to Requires=. However, if the units listed here are not started already, they will not be started and the starting of this unit will fail immediately.
     * 
     */
    public List<String> requisite() {
        return this.requisite == null ? List.of() : this.requisite;
    }
    /**
     * @return Configures (weak) requirement dependencies on other units.
     * 
     */
    public List<String> wants() {
        return this.wants == null ? List.of() : this.wants;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SystemdUnitSection defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable List<String> bindsTo;
        private @Nullable String description;
        private @Nullable List<String> documentation;
        private @Nullable List<String> requires;
        private @Nullable List<String> requisite;
        private @Nullable List<String> wants;
        public Builder() {}
        public Builder(SystemdUnitSection defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.bindsTo = defaults.bindsTo;
    	      this.description = defaults.description;
    	      this.documentation = defaults.documentation;
    	      this.requires = defaults.requires;
    	      this.requisite = defaults.requisite;
    	      this.wants = defaults.wants;
        }

        @CustomType.Setter
        public Builder bindsTo(@Nullable List<String> bindsTo) {

            this.bindsTo = bindsTo;
            return this;
        }
        public Builder bindsTo(String... bindsTo) {
            return bindsTo(List.of(bindsTo));
        }
        @CustomType.Setter
        public Builder description(@Nullable String description) {

            this.description = description;
            return this;
        }
        @CustomType.Setter
        public Builder documentation(@Nullable List<String> documentation) {

            this.documentation = documentation;
            return this;
        }
        public Builder documentation(String... documentation) {
            return documentation(List.of(documentation));
        }
        @CustomType.Setter
        public Builder requires(@Nullable List<String> requires) {

            this.requires = requires;
            return this;
        }
        public Builder requires(String... requires) {
            return requires(List.of(requires));
        }
        @CustomType.Setter
        public Builder requisite(@Nullable List<String> requisite) {

            this.requisite = requisite;
            return this;
        }
        public Builder requisite(String... requisite) {
            return requisite(List.of(requisite));
        }
        @CustomType.Setter
        public Builder wants(@Nullable List<String> wants) {

            this.wants = wants;
            return this;
        }
        public Builder wants(String... wants) {
            return wants(List.of(wants));
        }
        public SystemdUnitSection build() {
            final var _resultValue = new SystemdUnitSection();
            _resultValue.bindsTo = bindsTo;
            _resultValue.description = description;
            _resultValue.documentation = documentation;
            _resultValue.requires = requires;
            _resultValue.requisite = requisite;
            _resultValue.wants = wants;
            return _resultValue;
        }
    }
}
