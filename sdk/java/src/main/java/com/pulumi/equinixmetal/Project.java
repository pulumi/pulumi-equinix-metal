// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.equinixmetal;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.equinixmetal.ProjectArgs;
import com.pulumi.equinixmetal.Utilities;
import com.pulumi.equinixmetal.inputs.ProjectState;
import com.pulumi.equinixmetal.outputs.ProjectBgpConfig;
import java.lang.Boolean;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides an Equinix Metal project resource to allow you manage devices
 * in your projects.
 * 
 * &gt; Keep in mind that Equinix Metal invoicing is per project, so creating many `equinix-metal.Project` resources will affect the rendered invoice. If you want to keep your Equinix Metal bill simple and easy to review, please re-use your existing projects.
 * 
 * ## Example Usage
 * ### Create a new project
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.equinixmetal.Project;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var tfProject1 = new Project(&#34;tfProject1&#34;);
 * 
 *     }
 * }
 * ```
 * ### Example with BGP config
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.equinixmetal.Project;
 * import com.pulumi.equinixmetal.ProjectArgs;
 * import com.pulumi.equinixmetal.inputs.ProjectBgpConfigArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var tfProject1 = new Project(&#34;tfProject1&#34;, ProjectArgs.builder()        
 *             .bgpConfig(ProjectBgpConfigArgs.builder()
 *                 .asn(65000)
 *                 .deploymentType(&#34;local&#34;)
 *                 .md5(&#34;C179c28c41a85b&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * This resource can be imported using an existing project ID
 * 
 * ```sh
 *  $ pulumi import equinix-metal:index/project:Project metal_project {existing_project_id}
 * ```
 * 
 */
@ResourceType(type="equinix-metal:index/project:Project")
public class Project extends com.pulumi.resources.CustomResource {
    /**
     * Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
     * 
     */
    @Export(name="backendTransfer", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> backendTransfer;

    /**
     * @return Enable or disable [Backend Transfer](https://metal.equinix.com/developers/docs/networking/backend-transfer/), default is false
     * 
     */
    public Output<Optional<Boolean>> backendTransfer() {
        return Codegen.optional(this.backendTransfer);
    }
    /**
     * Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
     * 
     * Once you set the BGP config in a project, it can&#39;t be removed (due to a limitation in the Equinix Metal API). It can be updated.
     * 
     */
    @Export(name="bgpConfig", refs={ProjectBgpConfig.class}, tree="[0]")
    private Output</* @Nullable */ ProjectBgpConfig> bgpConfig;

    /**
     * @return Optional BGP settings. Refer to [Equinix Metal guide for BGP](https://metal.equinix.com/developers/docs/networking/local-global-bgp/).
     * 
     * Once you set the BGP config in a project, it can&#39;t be removed (due to a limitation in the Equinix Metal API). It can be updated.
     * 
     */
    public Output<Optional<ProjectBgpConfig>> bgpConfig() {
        return Codegen.optional(this.bgpConfig);
    }
    /**
     * The timestamp for when the project was created
     * 
     */
    @Export(name="created", refs={String.class}, tree="[0]")
    private Output<String> created;

    /**
     * @return The timestamp for when the project was created
     * 
     */
    public Output<String> created() {
        return this.created;
    }
    /**
     * The name of the project
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return The name of the project
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account.
     * 
     */
    @Export(name="organizationId", refs={String.class}, tree="[0]")
    private Output<String> organizationId;

    /**
     * @return The UUID of organization under which you want to create the project. If you leave it out, the project will be create under your the default organization of your account.
     * 
     */
    public Output<String> organizationId() {
        return this.organizationId;
    }
    /**
     * The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with `organization_id`, or default).
     * 
     */
    @Export(name="paymentMethodId", refs={String.class}, tree="[0]")
    private Output<String> paymentMethodId;

    /**
     * @return The UUID of payment method for this project. The payment method and the project need to belong to the same organization (passed with `organization_id`, or default).
     * 
     */
    public Output<String> paymentMethodId() {
        return this.paymentMethodId;
    }
    /**
     * The timestamp for the last time the project was updated
     * 
     */
    @Export(name="updated", refs={String.class}, tree="[0]")
    private Output<String> updated;

    /**
     * @return The timestamp for the last time the project was updated
     * 
     */
    public Output<String> updated() {
        return this.updated;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Project(String name) {
        this(name, ProjectArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Project(String name, @Nullable ProjectArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Project(String name, @Nullable ProjectArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix-metal:index/project:Project", name, args == null ? ProjectArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Project(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("equinix-metal:index/project:Project", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static Project get(String name, Output<String> id, @Nullable ProjectState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Project(name, id, state, options);
    }
}
