// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderArgs Empty = new ProviderArgs();

    /**
     * The API auth key for API operations.
     * 
     */
    @Import(name="authToken", required=true)
    private Output<String> authToken;

    /**
     * @return The API auth key for API operations.
     * 
     */
    public Output<String> authToken() {
        return this.authToken;
    }

    @Import(name="maxRetries", json=true)
    private @Nullable Output<Integer> maxRetries;

    public Optional<Output<Integer>> maxRetries() {
        return Optional.ofNullable(this.maxRetries);
    }

    @Import(name="maxRetryWaitSeconds", json=true)
    private @Nullable Output<Integer> maxRetryWaitSeconds;

    public Optional<Output<Integer>> maxRetryWaitSeconds() {
        return Optional.ofNullable(this.maxRetryWaitSeconds);
    }

    private ProviderArgs() {}

    private ProviderArgs(ProviderArgs $) {
        this.authToken = $.authToken;
        this.maxRetries = $.maxRetries;
        this.maxRetryWaitSeconds = $.maxRetryWaitSeconds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderArgs $;

        public Builder() {
            $ = new ProviderArgs();
        }

        public Builder(ProviderArgs defaults) {
            $ = new ProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param authToken The API auth key for API operations.
         * 
         * @return builder
         * 
         */
        public Builder authToken(Output<String> authToken) {
            $.authToken = authToken;
            return this;
        }

        /**
         * @param authToken The API auth key for API operations.
         * 
         * @return builder
         * 
         */
        public Builder authToken(String authToken) {
            return authToken(Output.of(authToken));
        }

        public Builder maxRetries(@Nullable Output<Integer> maxRetries) {
            $.maxRetries = maxRetries;
            return this;
        }

        public Builder maxRetries(Integer maxRetries) {
            return maxRetries(Output.of(maxRetries));
        }

        public Builder maxRetryWaitSeconds(@Nullable Output<Integer> maxRetryWaitSeconds) {
            $.maxRetryWaitSeconds = maxRetryWaitSeconds;
            return this;
        }

        public Builder maxRetryWaitSeconds(Integer maxRetryWaitSeconds) {
            return maxRetryWaitSeconds(Output.of(maxRetryWaitSeconds));
        }

        public ProviderArgs build() {
            $.authToken = Objects.requireNonNull($.authToken, "expected parameter 'authToken' to be non-null");
            return $;
        }
    }

}
