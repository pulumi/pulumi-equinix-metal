// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class GetVirtualCircuitArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetVirtualCircuitArgs Empty = new GetVirtualCircuitArgs();

    /**
     * ID of the virtual circuit resource
     * 
     */
    @Import(name="virtualCircuitId", required=true)
    private Output<String> virtualCircuitId;

    /**
     * @return ID of the virtual circuit resource
     * 
     */
    public Output<String> virtualCircuitId() {
        return this.virtualCircuitId;
    }

    private GetVirtualCircuitArgs() {}

    private GetVirtualCircuitArgs(GetVirtualCircuitArgs $) {
        this.virtualCircuitId = $.virtualCircuitId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetVirtualCircuitArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetVirtualCircuitArgs $;

        public Builder() {
            $ = new GetVirtualCircuitArgs();
        }

        public Builder(GetVirtualCircuitArgs defaults) {
            $ = new GetVirtualCircuitArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param virtualCircuitId ID of the virtual circuit resource
         * 
         * @return builder
         * 
         */
        public Builder virtualCircuitId(Output<String> virtualCircuitId) {
            $.virtualCircuitId = virtualCircuitId;
            return this;
        }

        /**
         * @param virtualCircuitId ID of the virtual circuit resource
         * 
         * @return builder
         * 
         */
        public Builder virtualCircuitId(String virtualCircuitId) {
            return virtualCircuitId(Output.of(virtualCircuitId));
        }

        public GetVirtualCircuitArgs build() {
            $.virtualCircuitId = Objects.requireNonNull($.virtualCircuitId, "expected parameter 'virtualCircuitId' to be non-null");
            return $;
        }
    }

}