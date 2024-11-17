# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from sso import permissions_pb2 as sso_dot_permissions__pb2

GRPC_GENERATED_VERSION = '1.68.0'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in sso/permissions_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class DutyPermissionsStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.CanRead = channel.unary_unary(
                '/permissions.DutyPermissions/CanRead',
                request_serializer=sso_dot_permissions__pb2.CanReadRequest.SerializeToString,
                response_deserializer=sso_dot_permissions__pb2.CanReadResponse.FromString,
                _registered_method=True)
        self.CanCreate = channel.unary_unary(
                '/permissions.DutyPermissions/CanCreate',
                request_serializer=sso_dot_permissions__pb2.CanCreateRequest.SerializeToString,
                response_deserializer=sso_dot_permissions__pb2.CanCreateResponse.FromString,
                _registered_method=True)
        self.CanUpdate = channel.unary_unary(
                '/permissions.DutyPermissions/CanUpdate',
                request_serializer=sso_dot_permissions__pb2.CanUpdateRequest.SerializeToString,
                response_deserializer=sso_dot_permissions__pb2.CanUpdateResponse.FromString,
                _registered_method=True)
        self.CanDelete = channel.unary_unary(
                '/permissions.DutyPermissions/CanDelete',
                request_serializer=sso_dot_permissions__pb2.CanDeleteRequest.SerializeToString,
                response_deserializer=sso_dot_permissions__pb2.CanDeleteResponse.FromString,
                _registered_method=True)


class DutyPermissionsServicer(object):
    """Missing associated documentation comment in .proto file."""

    def CanRead(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CanCreate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CanUpdate(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CanDelete(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DutyPermissionsServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'CanRead': grpc.unary_unary_rpc_method_handler(
                    servicer.CanRead,
                    request_deserializer=sso_dot_permissions__pb2.CanReadRequest.FromString,
                    response_serializer=sso_dot_permissions__pb2.CanReadResponse.SerializeToString,
            ),
            'CanCreate': grpc.unary_unary_rpc_method_handler(
                    servicer.CanCreate,
                    request_deserializer=sso_dot_permissions__pb2.CanCreateRequest.FromString,
                    response_serializer=sso_dot_permissions__pb2.CanCreateResponse.SerializeToString,
            ),
            'CanUpdate': grpc.unary_unary_rpc_method_handler(
                    servicer.CanUpdate,
                    request_deserializer=sso_dot_permissions__pb2.CanUpdateRequest.FromString,
                    response_serializer=sso_dot_permissions__pb2.CanUpdateResponse.SerializeToString,
            ),
            'CanDelete': grpc.unary_unary_rpc_method_handler(
                    servicer.CanDelete,
                    request_deserializer=sso_dot_permissions__pb2.CanDeleteRequest.FromString,
                    response_serializer=sso_dot_permissions__pb2.CanDeleteResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'permissions.DutyPermissions', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('permissions.DutyPermissions', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class DutyPermissions(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def CanRead(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/permissions.DutyPermissions/CanRead',
            sso_dot_permissions__pb2.CanReadRequest.SerializeToString,
            sso_dot_permissions__pb2.CanReadResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CanCreate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/permissions.DutyPermissions/CanCreate',
            sso_dot_permissions__pb2.CanCreateRequest.SerializeToString,
            sso_dot_permissions__pb2.CanCreateResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CanUpdate(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/permissions.DutyPermissions/CanUpdate',
            sso_dot_permissions__pb2.CanUpdateRequest.SerializeToString,
            sso_dot_permissions__pb2.CanUpdateResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def CanDelete(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/permissions.DutyPermissions/CanDelete',
            sso_dot_permissions__pb2.CanDeleteRequest.SerializeToString,
            sso_dot_permissions__pb2.CanDeleteResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)