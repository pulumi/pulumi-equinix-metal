// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.equinixmetal.inputs.ProjectBgpConfigArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProjectArgs Empty = new ProjectArgs();

    /**
     * Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
     * 
     */
    @Import(name="backendTransfer")
    private @Nullable Output<Boolean> backendTransfer;

    /**
     * @return Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
     * 
     */
    public Optional<Output<Boolean>> backendTransfer() {
        return Optional.ofNullable(this.backendTransfer);
    }

    /**
     * Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
     * 
     * Once you set the BGP config in a project, it can&#39;t be removed (due to a limitation in the Equinix Metal API). It can be updated.
     * 
     */
    @Import(name="bgpConfig")
    private @Nullable Output<ProjectBgpConfigArgs> bgpConfig;

    /**
     * @return Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
     * 
     * Once you set the BGP config in a project, it can&#39;t be removed (due to a limitation in the Equinix Metal API). It can be updated.
     * 
     */
    public Optional<Output<ProjectBgpConfigArgs>> bgpConfig() {
        return Optional.ofNullable(this.bgpConfig);
    }

    /**
     * The name of the project
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the project
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account.
     * 
     */
    @Import(name="organizationId")
    private @Nullable Output<String> organizationId;

    /**
     * @return The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account.
     * 
     */
    public Optional<Output<String>> organizationId() {
        return Optional.ofNullable(this.organizationId);
    }

    /**
     * The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with `organization_id`, or default).
     * 
     */
    @Import(name="paymentMethodId")
    private @Nullable Output<String> paymentMethodId;

    /**
     * @return The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with `organization_id`, or default).
     * 
     */
    public Optional<Output<String>> paymentMethodId() {
        return Optional.ofNullable(this.paymentMethodId);
    }

    private ProjectArgs() {}

    private ProjectArgs(ProjectArgs $) {
        this.backendTransfer = $.backendTransfer;
        this.bgpConfig = $.bgpConfig;
        this.name = $.name;
        this.organizationId = $.organizationId;
        this.paymentMethodId = $.paymentMethodId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectArgs $;

        public Builder() {
            $ = new ProjectArgs();
        }

        public Builder(ProjectArgs defaults) {
            $ = new ProjectArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backendTransfer Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
         * 
         * @return builder
         * 
         */
        public Builder backendTransfer(@Nullable Output<Boolean> backendTransfer) {
            $.backendTransfer = backendTransfer;
            return this;
        }

        /**
         * @param backendTransfer Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
         * 
         * @return builder
         * 
         */
        public Builder backendTransfer(Boolean backendTransfer) {
            return backendTransfer(Output.of(backendTransfer));
        }

        /**
         * @param bgpConfig Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
         * 
         * Once you set the BGP config in a project, it can&#39;t be removed (due to a limitation in the Equinix Metal API). It can be updated.
         * 
         * @return builder
         * 
         */
        public Builder bgpConfig(@Nullable Output<ProjectBgpConfigArgs> bgpConfig) {
            $.bgpConfig = bgpConfig;
            return this;
        }

        /**
         * @param bgpConfig Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
         * 
         * Once you set the BGP config in a project, it can&#39;t be removed (due to a limitation in the Equinix Metal API). It can be updated.
         * 
         * @return builder
         * 
         */
        public Builder bgpConfig(ProjectBgpConfigArgs bgpConfig) {
            return bgpConfig(Output.of(bgpConfig));
        }

        /**
         * @param name The name of the project
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the project
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param organizationId The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(@Nullable Output<String> organizationId) {
            $.organizationId = organizationId;
            return this;
        }

        /**
         * @param organizationId The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account.
         * 
         * @return builder
         * 
         */
        public Builder organizationId(String organizationId) {
            return organizationId(Output.of(organizationId));
        }

        /**
         * @param paymentMethodId The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with `organization_id`, or default).
         * 
         * @return builder
         * 
         */
        public Builder paymentMethodId(@Nullable Output<String> paymentMethodId) {
            $.paymentMethodId = paymentMethodId;
            return this;
        }

        /**
         * @param paymentMethodId The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with `organization_id`, or default).
         * 
         * @return builder
         * 
         */
        public Builder paymentMethodId(String paymentMethodId) {
            return paymentMethodId(Output.of(paymentMethodId));
        }

        public ProjectArgs build() {
            return $;
        }
    }

}
