// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.equinixmetal.inputs.SpotMarketRequestInstanceParametersArgs;
import java.lang.Boolean;
import java.lang.Double;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SpotMarketRequestArgs extends com.pulumi.resources.ResourceArgs {

    public static final SpotMarketRequestArgs Empty = new SpotMarketRequestArgs();

    /**
     * Maximum number devices to be created
     * 
     */
    @Import(name="devicesMax", required=true)
    private Output<Integer> devicesMax;

    /**
     * @return Maximum number devices to be created
     * 
     */
    public Output<Integer> devicesMax() {
        return this.devicesMax;
    }

    /**
     * Miniumum number devices to be created
     * 
     */
    @Import(name="devicesMin", required=true)
    private Output<Integer> devicesMin;

    /**
     * @return Miniumum number devices to be created
     * 
     */
    public Output<Integer> devicesMin() {
        return this.devicesMin;
    }

    /**
     * Facility IDs where devices should be created
     * 
     */
    @Import(name="facilities")
    private @Nullable Output<List<String>> facilities;

    /**
     * @return Facility IDs where devices should be created
     * 
     */
    public Optional<Output<List<String>>> facilities() {
        return Optional.ofNullable(this.facilities);
    }

    /**
     * Parameters for devices provisioned from this request. You can find the parameter description from the equinix-metal.Device doc.
     * * `billing_cycle`
     * * `plan`
     * * `operating_system`
     * * `hostname`
     * * `termintation_time`
     * * `always_pxe`
     * * `description`
     * * `features`
     * * `locked`
     * * `project_ssh_keys`
     * * `user_ssh_keys`
     * * `userdata`
     * * `customdata`
     * * `ipxe_script_url`
     * * `tags`
     * 
     */
    @Import(name="instanceParameters", required=true)
    private Output<SpotMarketRequestInstanceParametersArgs> instanceParameters;

    /**
     * @return Parameters for devices provisioned from this request. You can find the parameter description from the equinix-metal.Device doc.
     * * `billing_cycle`
     * * `plan`
     * * `operating_system`
     * * `hostname`
     * * `termintation_time`
     * * `always_pxe`
     * * `description`
     * * `features`
     * * `locked`
     * * `project_ssh_keys`
     * * `user_ssh_keys`
     * * `userdata`
     * * `customdata`
     * * `ipxe_script_url`
     * * `tags`
     * 
     */
    public Output<SpotMarketRequestInstanceParametersArgs> instanceParameters() {
        return this.instanceParameters;
    }

    /**
     * Maximum price user is willing to pay per hour per device
     * 
     */
    @Import(name="maxBidPrice", required=true)
    private Output<Double> maxBidPrice;

    /**
     * @return Maximum price user is willing to pay per hour per device
     * 
     */
    public Output<Double> maxBidPrice() {
        return this.maxBidPrice;
    }

    /**
     * Metro where devices should be created
     * 
     */
    @Import(name="metro")
    private @Nullable Output<String> metro;

    /**
     * @return Metro where devices should be created
     * 
     */
    public Optional<Output<String>> metro() {
        return Optional.ofNullable(this.metro);
    }

    /**
     * Project ID
     * 
     */
    @Import(name="projectId", required=true)
    private Output<String> projectId;

    /**
     * @return Project ID
     * 
     */
    public Output<String> projectId() {
        return this.projectId;
    }

    /**
     * On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
     * 
     */
    @Import(name="waitForDevices")
    private @Nullable Output<Boolean> waitForDevices;

    /**
     * @return On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
     * 
     */
    public Optional<Output<Boolean>> waitForDevices() {
        return Optional.ofNullable(this.waitForDevices);
    }

    private SpotMarketRequestArgs() {}

    private SpotMarketRequestArgs(SpotMarketRequestArgs $) {
        this.devicesMax = $.devicesMax;
        this.devicesMin = $.devicesMin;
        this.facilities = $.facilities;
        this.instanceParameters = $.instanceParameters;
        this.maxBidPrice = $.maxBidPrice;
        this.metro = $.metro;
        this.projectId = $.projectId;
        this.waitForDevices = $.waitForDevices;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SpotMarketRequestArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SpotMarketRequestArgs $;

        public Builder() {
            $ = new SpotMarketRequestArgs();
        }

        public Builder(SpotMarketRequestArgs defaults) {
            $ = new SpotMarketRequestArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param devicesMax Maximum number devices to be created
         * 
         * @return builder
         * 
         */
        public Builder devicesMax(Output<Integer> devicesMax) {
            $.devicesMax = devicesMax;
            return this;
        }

        /**
         * @param devicesMax Maximum number devices to be created
         * 
         * @return builder
         * 
         */
        public Builder devicesMax(Integer devicesMax) {
            return devicesMax(Output.of(devicesMax));
        }

        /**
         * @param devicesMin Miniumum number devices to be created
         * 
         * @return builder
         * 
         */
        public Builder devicesMin(Output<Integer> devicesMin) {
            $.devicesMin = devicesMin;
            return this;
        }

        /**
         * @param devicesMin Miniumum number devices to be created
         * 
         * @return builder
         * 
         */
        public Builder devicesMin(Integer devicesMin) {
            return devicesMin(Output.of(devicesMin));
        }

        /**
         * @param facilities Facility IDs where devices should be created
         * 
         * @return builder
         * 
         */
        public Builder facilities(@Nullable Output<List<String>> facilities) {
            $.facilities = facilities;
            return this;
        }

        /**
         * @param facilities Facility IDs where devices should be created
         * 
         * @return builder
         * 
         */
        public Builder facilities(List<String> facilities) {
            return facilities(Output.of(facilities));
        }

        /**
         * @param facilities Facility IDs where devices should be created
         * 
         * @return builder
         * 
         */
        public Builder facilities(String... facilities) {
            return facilities(List.of(facilities));
        }

        /**
         * @param instanceParameters Parameters for devices provisioned from this request. You can find the parameter description from the equinix-metal.Device doc.
         * * `billing_cycle`
         * * `plan`
         * * `operating_system`
         * * `hostname`
         * * `termintation_time`
         * * `always_pxe`
         * * `description`
         * * `features`
         * * `locked`
         * * `project_ssh_keys`
         * * `user_ssh_keys`
         * * `userdata`
         * * `customdata`
         * * `ipxe_script_url`
         * * `tags`
         * 
         * @return builder
         * 
         */
        public Builder instanceParameters(Output<SpotMarketRequestInstanceParametersArgs> instanceParameters) {
            $.instanceParameters = instanceParameters;
            return this;
        }

        /**
         * @param instanceParameters Parameters for devices provisioned from this request. You can find the parameter description from the equinix-metal.Device doc.
         * * `billing_cycle`
         * * `plan`
         * * `operating_system`
         * * `hostname`
         * * `termintation_time`
         * * `always_pxe`
         * * `description`
         * * `features`
         * * `locked`
         * * `project_ssh_keys`
         * * `user_ssh_keys`
         * * `userdata`
         * * `customdata`
         * * `ipxe_script_url`
         * * `tags`
         * 
         * @return builder
         * 
         */
        public Builder instanceParameters(SpotMarketRequestInstanceParametersArgs instanceParameters) {
            return instanceParameters(Output.of(instanceParameters));
        }

        /**
         * @param maxBidPrice Maximum price user is willing to pay per hour per device
         * 
         * @return builder
         * 
         */
        public Builder maxBidPrice(Output<Double> maxBidPrice) {
            $.maxBidPrice = maxBidPrice;
            return this;
        }

        /**
         * @param maxBidPrice Maximum price user is willing to pay per hour per device
         * 
         * @return builder
         * 
         */
        public Builder maxBidPrice(Double maxBidPrice) {
            return maxBidPrice(Output.of(maxBidPrice));
        }

        /**
         * @param metro Metro where devices should be created
         * 
         * @return builder
         * 
         */
        public Builder metro(@Nullable Output<String> metro) {
            $.metro = metro;
            return this;
        }

        /**
         * @param metro Metro where devices should be created
         * 
         * @return builder
         * 
         */
        public Builder metro(String metro) {
            return metro(Output.of(metro));
        }

        /**
         * @param projectId Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param waitForDevices On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
         * 
         * @return builder
         * 
         */
        public Builder waitForDevices(@Nullable Output<Boolean> waitForDevices) {
            $.waitForDevices = waitForDevices;
            return this;
        }

        /**
         * @param waitForDevices On resource creation - wait until all desired devices are active, on resource destruction - wait until devices are removed
         * 
         * @return builder
         * 
         */
        public Builder waitForDevices(Boolean waitForDevices) {
            return waitForDevices(Output.of(waitForDevices));
        }

        public SpotMarketRequestArgs build() {
            $.devicesMax = Objects.requireNonNull($.devicesMax, "expected parameter 'devicesMax' to be non-null");
            $.devicesMin = Objects.requireNonNull($.devicesMin, "expected parameter 'devicesMin' to be non-null");
            $.instanceParameters = Objects.requireNonNull($.instanceParameters, "expected parameter 'instanceParameters' to be non-null");
            $.maxBidPrice = Objects.requireNonNull($.maxBidPrice, "expected parameter 'maxBidPrice' to be non-null");
            $.projectId = Objects.requireNonNull($.projectId, "expected parameter 'projectId' to be non-null");
            return $;
        }
    }

}
