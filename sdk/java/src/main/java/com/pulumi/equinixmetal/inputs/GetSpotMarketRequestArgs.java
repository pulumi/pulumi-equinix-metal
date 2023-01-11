// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetSpotMarketRequestArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSpotMarketRequestArgs Empty = new GetSpotMarketRequestArgs();

    /**
     * The id of the Spot Market Request
     * 
     */
    @Import(name="requestId", required=true)
    private Output<String> requestId;

    /**
     * @return The id of the Spot Market Request
     * 
     */
    public Output<String> requestId() {
        return this.requestId;
    }

    private GetSpotMarketRequestArgs() {}

    private GetSpotMarketRequestArgs(GetSpotMarketRequestArgs $) {
        this.requestId = $.requestId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSpotMarketRequestArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSpotMarketRequestArgs $;

        public Builder() {
            $ = new GetSpotMarketRequestArgs();
        }

        public Builder(GetSpotMarketRequestArgs defaults) {
            $ = new GetSpotMarketRequestArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param requestId The id of the Spot Market Request
         * 
         * @return builder
         * 
         */
        public Builder requestId(Output<String> requestId) {
            $.requestId = requestId;
            return this;
        }

        /**
         * @param requestId The id of the Spot Market Request
         * 
         * @return builder
         * 
         */
        public Builder requestId(String requestId) {
            return requestId(Output.of(requestId));
        }

        public GetSpotMarketRequestArgs build() {
            $.requestId = Objects.requireNonNull($.requestId, "expected parameter 'requestId' to be non-null");
            return $;
        }
    }

}