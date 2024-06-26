// *** WARNING: this file was generated by pulumi. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.unmango.kubernetesthehardway.tools.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


/**
 * Abstraction over the `tar` utility on a remote system.
 * 
 */
public final class TarOptsArgs extends com.pulumi.resources.ResourceArgs {

    public static final TarOptsArgs Empty = new TarOptsArgs();

    /**
     * Corresponds to the [ARCHIVE] argument.
     * 
     */
    @Import(name="archive", required=true)
    private Output<String> archive;

    /**
     * @return Corresponds to the [ARCHIVE] argument.
     * 
     */
    public Output<String> archive() {
        return this.archive;
    }

    /**
     * Corresponds to the `--directory` option.
     * 
     */
    @Import(name="directory")
    private @Nullable Output<String> directory;

    /**
     * @return Corresponds to the `--directory` option.
     * 
     */
    public Optional<Output<String>> directory() {
        return Optional.ofNullable(this.directory);
    }

    /**
     * Corresponds to the `--extract` option.
     * 
     */
    @Import(name="extract")
    private @Nullable Output<Boolean> extract;

    /**
     * @return Corresponds to the `--extract` option.
     * 
     */
    public Optional<Output<Boolean>> extract() {
        return Optional.ofNullable(this.extract);
    }

    /**
     * Corresponds to the [FILE] argument.
     * 
     */
    @Import(name="files")
    private @Nullable Output<List<String>> files;

    /**
     * @return Corresponds to the [FILE] argument.
     * 
     */
    public Optional<Output<List<String>>> files() {
        return Optional.ofNullable(this.files);
    }

    /**
     * Corresponds to the `--gzip` option.
     * 
     */
    @Import(name="gzip")
    private @Nullable Output<Boolean> gzip;

    /**
     * @return Corresponds to the `--gzip` option.
     * 
     */
    public Optional<Output<Boolean>> gzip() {
        return Optional.ofNullable(this.gzip);
    }

    /**
     * Whether rm should be run when the resource is created or deleted.
     * 
     */
    @Import(name="onDelete")
    private @Nullable Output<Boolean> onDelete;

    /**
     * @return Whether rm should be run when the resource is created or deleted.
     * 
     */
    public Optional<Output<Boolean>> onDelete() {
        return Optional.ofNullable(this.onDelete);
    }

    /**
     * Corresponds to the `--recursive` option.
     * 
     */
    @Import(name="recursive")
    private @Nullable Output<Boolean> recursive;

    /**
     * @return Corresponds to the `--recursive` option.
     * 
     */
    public Optional<Output<Boolean>> recursive() {
        return Optional.ofNullable(this.recursive);
    }

    /**
     * Corresponds to the `--strip-components` option.
     * 
     */
    @Import(name="stripComponents")
    private @Nullable Output<Integer> stripComponents;

    /**
     * @return Corresponds to the `--strip-components` option.
     * 
     */
    public Optional<Output<Integer>> stripComponents() {
        return Optional.ofNullable(this.stripComponents);
    }

    private TarOptsArgs() {}

    private TarOptsArgs(TarOptsArgs $) {
        this.archive = $.archive;
        this.directory = $.directory;
        this.extract = $.extract;
        this.files = $.files;
        this.gzip = $.gzip;
        this.onDelete = $.onDelete;
        this.recursive = $.recursive;
        this.stripComponents = $.stripComponents;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(TarOptsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private TarOptsArgs $;

        public Builder() {
            $ = new TarOptsArgs();
        }

        public Builder(TarOptsArgs defaults) {
            $ = new TarOptsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param archive Corresponds to the [ARCHIVE] argument.
         * 
         * @return builder
         * 
         */
        public Builder archive(Output<String> archive) {
            $.archive = archive;
            return this;
        }

        /**
         * @param archive Corresponds to the [ARCHIVE] argument.
         * 
         * @return builder
         * 
         */
        public Builder archive(String archive) {
            return archive(Output.of(archive));
        }

        /**
         * @param directory Corresponds to the `--directory` option.
         * 
         * @return builder
         * 
         */
        public Builder directory(@Nullable Output<String> directory) {
            $.directory = directory;
            return this;
        }

        /**
         * @param directory Corresponds to the `--directory` option.
         * 
         * @return builder
         * 
         */
        public Builder directory(String directory) {
            return directory(Output.of(directory));
        }

        /**
         * @param extract Corresponds to the `--extract` option.
         * 
         * @return builder
         * 
         */
        public Builder extract(@Nullable Output<Boolean> extract) {
            $.extract = extract;
            return this;
        }

        /**
         * @param extract Corresponds to the `--extract` option.
         * 
         * @return builder
         * 
         */
        public Builder extract(Boolean extract) {
            return extract(Output.of(extract));
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(@Nullable Output<List<String>> files) {
            $.files = files;
            return this;
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(List<String> files) {
            return files(Output.of(files));
        }

        /**
         * @param files Corresponds to the [FILE] argument.
         * 
         * @return builder
         * 
         */
        public Builder files(String... files) {
            return files(List.of(files));
        }

        /**
         * @param gzip Corresponds to the `--gzip` option.
         * 
         * @return builder
         * 
         */
        public Builder gzip(@Nullable Output<Boolean> gzip) {
            $.gzip = gzip;
            return this;
        }

        /**
         * @param gzip Corresponds to the `--gzip` option.
         * 
         * @return builder
         * 
         */
        public Builder gzip(Boolean gzip) {
            return gzip(Output.of(gzip));
        }

        /**
         * @param onDelete Whether rm should be run when the resource is created or deleted.
         * 
         * @return builder
         * 
         */
        public Builder onDelete(@Nullable Output<Boolean> onDelete) {
            $.onDelete = onDelete;
            return this;
        }

        /**
         * @param onDelete Whether rm should be run when the resource is created or deleted.
         * 
         * @return builder
         * 
         */
        public Builder onDelete(Boolean onDelete) {
            return onDelete(Output.of(onDelete));
        }

        /**
         * @param recursive Corresponds to the `--recursive` option.
         * 
         * @return builder
         * 
         */
        public Builder recursive(@Nullable Output<Boolean> recursive) {
            $.recursive = recursive;
            return this;
        }

        /**
         * @param recursive Corresponds to the `--recursive` option.
         * 
         * @return builder
         * 
         */
        public Builder recursive(Boolean recursive) {
            return recursive(Output.of(recursive));
        }

        /**
         * @param stripComponents Corresponds to the `--strip-components` option.
         * 
         * @return builder
         * 
         */
        public Builder stripComponents(@Nullable Output<Integer> stripComponents) {
            $.stripComponents = stripComponents;
            return this;
        }

        /**
         * @param stripComponents Corresponds to the `--strip-components` option.
         * 
         * @return builder
         * 
         */
        public Builder stripComponents(Integer stripComponents) {
            return stripComponents(Output.of(stripComponents));
        }

        public TarOptsArgs build() {
            if ($.archive == null) {
                throw new MissingRequiredPropertyException("TarOptsArgs", "archive");
            }
            return $;
        }
    }

}
