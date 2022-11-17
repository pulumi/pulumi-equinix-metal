// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSpotMarketPricePlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSpotMarketPricePlainArgs Empty = new GetSpotMarketPricePlainArgs();

    /**
     * Name of the facility.
     * 
     */
    @Import(name="facility")
    private @Nullable String facility;

    /**
     * @return Name of the facility.
     * 
     */
    public Optional<String> facility() {
        return Optional.ofNullable(this.facility);
    }

    /**
     * Name of the metro.
     * 
     */
    @Import(name="metro")
    private @Nullable String metro;

    /**
     * @return Name of the metro.
     * 
     */
    public Optional<String> metro() {
        return Optional.ofNullable(this.metro);
    }

    /**
     * Name of the plan.
     * 
     */
    @Import(name="plan", required=true)
    private String plan;

    /**
     * @return Name of the plan.
     * 
     */
    public String plan() {
        return this.plan;
    }

    private GetSpotMarketPricePlainArgs() {}

    private GetSpotMarketPricePlainArgs(GetSpotMarketPricePlainArgs $) {
        this.facility = $.facility;
        this.metro = $.metro;
        this.plan = $.plan;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSpotMarketPricePlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSpotMarketPricePlainArgs $;

        public Builder() {
            $ = new GetSpotMarketPricePlainArgs();
        }

        public Builder(GetSpotMarketPricePlainArgs defaults) {
            $ = new GetSpotMarketPricePlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param facility Name of the facility.
         * 
         * @return builder
         * 
         */
        public Builder facility(@Nullable String facility) {
            $.facility = facility;
            return this;
        }

        /**
         * @param metro Name of the metro.
         * 
         * @return builder
         * 
         */
        public Builder metro(@Nullable String metro) {
            $.metro = metro;
            return this;
        }

        /**
         * @param plan Name of the plan.
         * 
         * @return builder
         * 
         */
        public Builder plan(String plan) {
            $.plan = plan;
            return this;
        }

        public GetSpotMarketPricePlainArgs build() {
            $.plan = Objects.requireNonNull($.plan, "expected parameter 'plan' to be non-null");
            return $;
        }
    }

}
