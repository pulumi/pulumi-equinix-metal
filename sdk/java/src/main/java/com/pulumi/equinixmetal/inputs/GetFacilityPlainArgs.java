// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.equinixmetal.inputs.GetFacilityCapacity;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetFacilityPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetFacilityPlainArgs Empty = new GetFacilityPlainArgs();

    /**
     * (Optional) Ensure that queried facility has capacity for specified number of given plans
     * 
     */
    @Import(name="capacities")
    private @Nullable List<GetFacilityCapacity> capacities;

    /**
     * @return (Optional) Ensure that queried facility has capacity for specified number of given plans
     * 
     */
    public Optional<List<GetFacilityCapacity>> capacities() {
        return Optional.ofNullable(this.capacities);
    }

    /**
     * The facility code
     * 
     */
    @Import(name="code", required=true)
    private String code;

    /**
     * @return The facility code
     * 
     */
    public String code() {
        return this.code;
    }

    /**
     * Set of feature strings that the facility must have
     * 
     * Facilities can be looked up by `code`.
     * 
     */
    @Import(name="featuresRequireds")
    private @Nullable List<String> featuresRequireds;

    /**
     * @return Set of feature strings that the facility must have
     * 
     * Facilities can be looked up by `code`.
     * 
     */
    public Optional<List<String>> featuresRequireds() {
        return Optional.ofNullable(this.featuresRequireds);
    }

    private GetFacilityPlainArgs() {}

    private GetFacilityPlainArgs(GetFacilityPlainArgs $) {
        this.capacities = $.capacities;
        this.code = $.code;
        this.featuresRequireds = $.featuresRequireds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetFacilityPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetFacilityPlainArgs $;

        public Builder() {
            $ = new GetFacilityPlainArgs();
        }

        public Builder(GetFacilityPlainArgs defaults) {
            $ = new GetFacilityPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param capacities (Optional) Ensure that queried facility has capacity for specified number of given plans
         * 
         * @return builder
         * 
         */
        public Builder capacities(@Nullable List<GetFacilityCapacity> capacities) {
            $.capacities = capacities;
            return this;
        }

        /**
         * @param capacities (Optional) Ensure that queried facility has capacity for specified number of given plans
         * 
         * @return builder
         * 
         */
        public Builder capacities(GetFacilityCapacity... capacities) {
            return capacities(List.of(capacities));
        }

        /**
         * @param code The facility code
         * 
         * @return builder
         * 
         */
        public Builder code(String code) {
            $.code = code;
            return this;
        }

        /**
         * @param featuresRequireds Set of feature strings that the facility must have
         * 
         * Facilities can be looked up by `code`.
         * 
         * @return builder
         * 
         */
        public Builder featuresRequireds(@Nullable List<String> featuresRequireds) {
            $.featuresRequireds = featuresRequireds;
            return this;
        }

        /**
         * @param featuresRequireds Set of feature strings that the facility must have
         * 
         * Facilities can be looked up by `code`.
         * 
         * @return builder
         * 
         */
        public Builder featuresRequireds(String... featuresRequireds) {
            return featuresRequireds(List.of(featuresRequireds));
        }

        public GetFacilityPlainArgs build() {
            $.code = Objects.requireNonNull($.code, "expected parameter 'code' to be non-null");
            return $;
        }
    }

}
