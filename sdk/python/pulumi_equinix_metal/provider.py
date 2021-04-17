# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 auth_token: pulumi.Input[str],
                 max_retries: Optional[pulumi.Input[int]] = None,
                 max_retry_wait_seconds: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a Provider resource.
        :param pulumi.Input[str] auth_token: The API auth key for API operations.
        """
        pulumi.set(__self__, "auth_token", auth_token)
        if max_retries is not None:
            pulumi.set(__self__, "max_retries", max_retries)
        if max_retry_wait_seconds is not None:
            pulumi.set(__self__, "max_retry_wait_seconds", max_retry_wait_seconds)

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> pulumi.Input[str]:
        """
        The API auth key for API operations.
        """
        return pulumi.get(self, "auth_token")

    @auth_token.setter
    def auth_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "auth_token", value)

    @property
    @pulumi.getter(name="maxRetries")
    def max_retries(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_retries")

    @max_retries.setter
    def max_retries(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retries", value)

    @property
    @pulumi.getter(name="maxRetryWaitSeconds")
    def max_retry_wait_seconds(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "max_retry_wait_seconds")

    @max_retry_wait_seconds.setter
    def max_retry_wait_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_retry_wait_seconds", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 max_retry_wait_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        The provider type for the metal package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] auth_token: The API auth key for API operations.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        The provider type for the metal package. By default, resources use package-wide configuration
        settings, however an explicit `Provider` instance may be created and passed during resource
        construction to achieve fine-grained programmatic control over provider settings. See the
        [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.

        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 max_retries: Optional[pulumi.Input[int]] = None,
                 max_retry_wait_seconds: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if auth_token is None and not opts.urn:
                raise TypeError("Missing required property 'auth_token'")
            __props__.__dict__["auth_token"] = auth_token
            __props__.__dict__["max_retries"] = pulumi.Output.from_input(max_retries).apply(pulumi.runtime.to_json) if max_retries is not None else None
            __props__.__dict__["max_retry_wait_seconds"] = pulumi.Output.from_input(max_retry_wait_seconds).apply(pulumi.runtime.to_json) if max_retry_wait_seconds is not None else None
        super(Provider, __self__).__init__(
            'equinix-metal',
            resource_name,
            __props__,
            opts)

