// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetOperatingSystemPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetOperatingSystemPlainArgs Empty = new GetOperatingSystemPlainArgs();

    /**
     * Name of the OS distribution.
     * 
     */
    @Import(name="distro")
    private @Nullable String distro;

    /**
     * @return Name of the OS distribution.
     * 
     */
    public Optional<String> distro() {
        return Optional.ofNullable(this.distro);
    }

    /**
     * Name or part of the name of the distribution. Case insensitive.
     * 
     */
    @Import(name="name")
    private @Nullable String name;

    /**
     * @return Name or part of the name of the distribution. Case insensitive.
     * 
     */
    public Optional<String> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Plan name.
     * 
     */
    @Import(name="provisionableOn")
    private @Nullable String provisionableOn;

    /**
     * @return Plan name.
     * 
     */
    public Optional<String> provisionableOn() {
        return Optional.ofNullable(this.provisionableOn);
    }

    /**
     * Version of the distribution
     * 
     */
    @Import(name="version")
    private @Nullable String version;

    /**
     * @return Version of the distribution
     * 
     */
    public Optional<String> version() {
        return Optional.ofNullable(this.version);
    }

    private GetOperatingSystemPlainArgs() {}

    private GetOperatingSystemPlainArgs(GetOperatingSystemPlainArgs $) {
        this.distro = $.distro;
        this.name = $.name;
        this.provisionableOn = $.provisionableOn;
        this.version = $.version;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetOperatingSystemPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetOperatingSystemPlainArgs $;

        public Builder() {
            $ = new GetOperatingSystemPlainArgs();
        }

        public Builder(GetOperatingSystemPlainArgs defaults) {
            $ = new GetOperatingSystemPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param distro Name of the OS distribution.
         * 
         * @return builder
         * 
         */
        public Builder distro(@Nullable String distro) {
            $.distro = distro;
            return this;
        }

        /**
         * @param name Name or part of the name of the distribution. Case insensitive.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable String name) {
            $.name = name;
            return this;
        }

        /**
         * @param provisionableOn Plan name.
         * 
         * @return builder
         * 
         */
        public Builder provisionableOn(@Nullable String provisionableOn) {
            $.provisionableOn = provisionableOn;
            return this;
        }

        /**
         * @param version Version of the distribution
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable String version) {
            $.version = version;
            return this;
        }

        public GetOperatingSystemPlainArgs build() {
            return $;
        }
    }

}
