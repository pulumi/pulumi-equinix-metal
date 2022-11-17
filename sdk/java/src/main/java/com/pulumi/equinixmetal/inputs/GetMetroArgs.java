// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.equinixmetal.inputs.GetMetroCapacityArgs;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetMetroArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetMetroArgs Empty = new GetMetroArgs();

    /**
     * (Optional) Ensure that queried metro has capacity for specified number of given plans
     * 
     */
    @Import(name="capacities")
    private @Nullable Output<List<GetMetroCapacityArgs>> capacities;

    /**
     * @return (Optional) Ensure that queried metro has capacity for specified number of given plans
     * 
     */
    public Optional<Output<List<GetMetroCapacityArgs>>> capacities() {
        return Optional.ofNullable(this.capacities);
    }

    /**
     * The metro code
     * 
     */
    @Import(name="code", required=true)
    private Output<String> code;

    /**
     * @return The metro code
     * 
     */
    public Output<String> code() {
        return this.code;
    }

    private GetMetroArgs() {}

    private GetMetroArgs(GetMetroArgs $) {
        this.capacities = $.capacities;
        this.code = $.code;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetMetroArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetMetroArgs $;

        public Builder() {
            $ = new GetMetroArgs();
        }

        public Builder(GetMetroArgs defaults) {
            $ = new GetMetroArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param capacities (Optional) Ensure that queried metro has capacity for specified number of given plans
         * 
         * @return builder
         * 
         */
        public Builder capacities(@Nullable Output<List<GetMetroCapacityArgs>> capacities) {
            $.capacities = capacities;
            return this;
        }

        /**
         * @param capacities (Optional) Ensure that queried metro has capacity for specified number of given plans
         * 
         * @return builder
         * 
         */
        public Builder capacities(List<GetMetroCapacityArgs> capacities) {
            return capacities(Output.of(capacities));
        }

        /**
         * @param capacities (Optional) Ensure that queried metro has capacity for specified number of given plans
         * 
         * @return builder
         * 
         */
        public Builder capacities(GetMetroCapacityArgs... capacities) {
            return capacities(List.of(capacities));
        }

        /**
         * @param code The metro code
         * 
         * @return builder
         * 
         */
        public Builder code(Output<String> code) {
            $.code = code;
            return this;
        }

        /**
         * @param code The metro code
         * 
         * @return builder
         * 
         */
        public Builder code(String code) {
            return code(Output.of(code));
        }

        public GetMetroArgs build() {
            $.code = Objects.requireNonNull($.code, "expected parameter 'code' to be non-null");
            return $;
        }
    }

}
