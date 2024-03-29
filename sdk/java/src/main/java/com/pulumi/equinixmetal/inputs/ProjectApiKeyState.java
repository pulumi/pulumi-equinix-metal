// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProjectApiKeyState extends com.pulumi.resources.ResourceArgs {

    public static final ProjectApiKeyState Empty = new ProjectApiKeyState();

    /**
     * Description string for the Project API Key resource
     * * `read-only` - Flag indicating whether the API key shoud be read-only
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description string for the Project API Key resource
     * * `read-only` - Flag indicating whether the API key shoud be read-only
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * UUID of the project where the API key is scoped to
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return UUID of the project where the API key is scoped to
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Flag indicating whether the API key shoud be read-only
     * 
     */
    @Import(name="readOnly")
    private @Nullable Output<Boolean> readOnly;

    /**
     * @return Flag indicating whether the API key shoud be read-only
     * 
     */
    public Optional<Output<Boolean>> readOnly() {
        return Optional.ofNullable(this.readOnly);
    }

    /**
     * API token which can be used in Equinix Metal API clients
     * 
     */
    @Import(name="token")
    private @Nullable Output<String> token;

    /**
     * @return API token which can be used in Equinix Metal API clients
     * 
     */
    public Optional<Output<String>> token() {
        return Optional.ofNullable(this.token);
    }

    private ProjectApiKeyState() {}

    private ProjectApiKeyState(ProjectApiKeyState $) {
        this.description = $.description;
        this.projectId = $.projectId;
        this.readOnly = $.readOnly;
        this.token = $.token;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProjectApiKeyState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProjectApiKeyState $;

        public Builder() {
            $ = new ProjectApiKeyState();
        }

        public Builder(ProjectApiKeyState defaults) {
            $ = new ProjectApiKeyState(Objects.requireNonNull(defaults));
        }

        /**
         * @param description Description string for the Project API Key resource
         * * `read-only` - Flag indicating whether the API key shoud be read-only
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description string for the Project API Key resource
         * * `read-only` - Flag indicating whether the API key shoud be read-only
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param projectId UUID of the project where the API key is scoped to
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId UUID of the project where the API key is scoped to
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param readOnly Flag indicating whether the API key shoud be read-only
         * 
         * @return builder
         * 
         */
        public Builder readOnly(@Nullable Output<Boolean> readOnly) {
            $.readOnly = readOnly;
            return this;
        }

        /**
         * @param readOnly Flag indicating whether the API key shoud be read-only
         * 
         * @return builder
         * 
         */
        public Builder readOnly(Boolean readOnly) {
            return readOnly(Output.of(readOnly));
        }

        /**
         * @param token API token which can be used in Equinix Metal API clients
         * 
         * @return builder
         * 
         */
        public Builder token(@Nullable Output<String> token) {
            $.token = token;
            return this;
        }

        /**
         * @param token API token which can be used in Equinix Metal API clients
         * 
         * @return builder
         * 
         */
        public Builder token(String token) {
            return token(Output.of(token));
        }

        public ProjectApiKeyState build() {
            return $;
        }
    }

}
