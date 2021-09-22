# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['UserApiKeyArgs', 'UserApiKey']

@pulumi.input_type
class UserApiKeyArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 read_only: pulumi.Input[bool]):
        """
        The set of arguments for constructing a UserApiKey resource.
        :param pulumi.Input[str] description: Description string for the User API Key resource
               * `read-only` - Flag indicating whether the API key shoud be read-only
        :param pulumi.Input[bool] read_only: Flag indicating whether the API key shoud be read-only
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "read_only", read_only)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        Description string for the User API Key resource
        * `read-only` - Flag indicating whether the API key shoud be read-only
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> pulumi.Input[bool]:
        """
        Flag indicating whether the API key shoud be read-only
        """
        return pulumi.get(self, "read_only")

    @read_only.setter
    def read_only(self, value: pulumi.Input[bool]):
        pulumi.set(self, "read_only", value)


@pulumi.input_type
class _UserApiKeyState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 read_only: Optional[pulumi.Input[bool]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UserApiKey resources.
        :param pulumi.Input[str] description: Description string for the User API Key resource
               * `read-only` - Flag indicating whether the API key shoud be read-only
        :param pulumi.Input[bool] read_only: Flag indicating whether the API key shoud be read-only
        :param pulumi.Input[str] token: API token which can be used in Equinix Metal API clients
        :param pulumi.Input[str] user_id: UUID of the owner of the API key
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if read_only is not None:
            pulumi.set(__self__, "read_only", read_only)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if user_id is not None:
            pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description string for the User API Key resource
        * `read-only` - Flag indicating whether the API key shoud be read-only
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag indicating whether the API key shoud be read-only
        """
        return pulumi.get(self, "read_only")

    @read_only.setter
    def read_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "read_only", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        API token which can be used in Equinix Metal API clients
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[pulumi.Input[str]]:
        """
        UUID of the owner of the API key
        """
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user_id", value)


class UserApiKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 read_only: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        test = equinix_metal.UserApiKey("test",
            description="Read-only user key",
            read_only=True)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description string for the User API Key resource
               * `read-only` - Flag indicating whether the API key shoud be read-only
        :param pulumi.Input[bool] read_only: Flag indicating whether the API key shoud be read-only
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserApiKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_equinix_metal as equinix_metal

        test = equinix_metal.UserApiKey("test",
            description="Read-only user key",
            read_only=True)
        ```

        :param str resource_name: The name of the resource.
        :param UserApiKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserApiKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 read_only: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserApiKeyArgs.__new__(UserApiKeyArgs)

            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if read_only is None and not opts.urn:
                raise TypeError("Missing required property 'read_only'")
            __props__.__dict__["read_only"] = read_only
            __props__.__dict__["token"] = None
            __props__.__dict__["user_id"] = None
        super(UserApiKey, __self__).__init__(
            'equinix-metal:index/userApiKey:UserApiKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            read_only: Optional[pulumi.Input[bool]] = None,
            token: Optional[pulumi.Input[str]] = None,
            user_id: Optional[pulumi.Input[str]] = None) -> 'UserApiKey':
        """
        Get an existing UserApiKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description string for the User API Key resource
               * `read-only` - Flag indicating whether the API key shoud be read-only
        :param pulumi.Input[bool] read_only: Flag indicating whether the API key shoud be read-only
        :param pulumi.Input[str] token: API token which can be used in Equinix Metal API clients
        :param pulumi.Input[str] user_id: UUID of the owner of the API key
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserApiKeyState.__new__(_UserApiKeyState)

        __props__.__dict__["description"] = description
        __props__.__dict__["read_only"] = read_only
        __props__.__dict__["token"] = token
        __props__.__dict__["user_id"] = user_id
        return UserApiKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Description string for the User API Key resource
        * `read-only` - Flag indicating whether the API key shoud be read-only
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> pulumi.Output[bool]:
        """
        Flag indicating whether the API key shoud be read-only
        """
        return pulumi.get(self, "read_only")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        API token which can be used in Equinix Metal API clients
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Output[str]:
        """
        UUID of the owner of the API key
        """
        return pulumi.get(self, "user_id")

