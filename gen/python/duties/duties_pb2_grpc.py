# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

from duties import duties_pb2 as duties_dot_duties__pb2

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
        + f' but the generated code in duties/duties_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
    )


class DutiesStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.ReadDuty = channel.unary_unary(
                '/duties.Duties/ReadDuty',
                request_serializer=duties_dot_duties__pb2.ReadDutyRequest.SerializeToString,
                response_deserializer=duties_dot_duties__pb2.ReadDutyResponse.FromString,
                _registered_method=True)
        self.ReadOwnerDuties = channel.unary_unary(
                '/duties.Duties/ReadOwnerDuties',
                request_serializer=duties_dot_duties__pb2.ReadOwnerDutiesRequest.SerializeToString,
                response_deserializer=duties_dot_duties__pb2.ReadDutiesResponse.FromString,
                _registered_method=True)
        self.ReadTargetDuties = channel.unary_unary(
                '/duties.Duties/ReadTargetDuties',
                request_serializer=duties_dot_duties__pb2.ReadTargetDutiesRequest.SerializeToString,
                response_deserializer=duties_dot_duties__pb2.ReadDutiesResponse.FromString,
                _registered_method=True)
        self.CreateDuty = channel.unary_unary(
                '/duties.Duties/CreateDuty',
                request_serializer=duties_dot_duties__pb2.CreateDutyRequest.SerializeToString,
                response_deserializer=duties_dot_duties__pb2.CreateDutyResponse.FromString,
                _registered_method=True)
        self.UpdateDuty = channel.unary_unary(
                '/duties.Duties/UpdateDuty',
                request_serializer=duties_dot_duties__pb2.UpdateDutyRequest.SerializeToString,
                response_deserializer=duties_dot_duties__pb2.UpdateDutyResponse.FromString,
                _registered_method=True)
        self.DeleteDuty = channel.unary_unary(
                '/duties.Duties/DeleteDuty',
                request_serializer=duties_dot_duties__pb2.DeleteDutyRequest.SerializeToString,
                response_deserializer=duties_dot_duties__pb2.DeleteDutyResponse.FromString,
                _registered_method=True)


class DutiesServicer(object):
    """Missing associated documentation comment in .proto file."""

    def ReadDuty(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ReadOwnerDuties(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def ReadTargetDuties(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def CreateDuty(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateDuty(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteDuty(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_DutiesServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'ReadDuty': grpc.unary_unary_rpc_method_handler(
                    servicer.ReadDuty,
                    request_deserializer=duties_dot_duties__pb2.ReadDutyRequest.FromString,
                    response_serializer=duties_dot_duties__pb2.ReadDutyResponse.SerializeToString,
            ),
            'ReadOwnerDuties': grpc.unary_unary_rpc_method_handler(
                    servicer.ReadOwnerDuties,
                    request_deserializer=duties_dot_duties__pb2.ReadOwnerDutiesRequest.FromString,
                    response_serializer=duties_dot_duties__pb2.ReadDutiesResponse.SerializeToString,
            ),
            'ReadTargetDuties': grpc.unary_unary_rpc_method_handler(
                    servicer.ReadTargetDuties,
                    request_deserializer=duties_dot_duties__pb2.ReadTargetDutiesRequest.FromString,
                    response_serializer=duties_dot_duties__pb2.ReadDutiesResponse.SerializeToString,
            ),
            'CreateDuty': grpc.unary_unary_rpc_method_handler(
                    servicer.CreateDuty,
                    request_deserializer=duties_dot_duties__pb2.CreateDutyRequest.FromString,
                    response_serializer=duties_dot_duties__pb2.CreateDutyResponse.SerializeToString,
            ),
            'UpdateDuty': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateDuty,
                    request_deserializer=duties_dot_duties__pb2.UpdateDutyRequest.FromString,
                    response_serializer=duties_dot_duties__pb2.UpdateDutyResponse.SerializeToString,
            ),
            'DeleteDuty': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteDuty,
                    request_deserializer=duties_dot_duties__pb2.DeleteDutyRequest.FromString,
                    response_serializer=duties_dot_duties__pb2.DeleteDutyResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'duties.Duties', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('duties.Duties', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class Duties(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def ReadDuty(request,
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
            '/duties.Duties/ReadDuty',
            duties_dot_duties__pb2.ReadDutyRequest.SerializeToString,
            duties_dot_duties__pb2.ReadDutyResponse.FromString,
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
    def ReadOwnerDuties(request,
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
            '/duties.Duties/ReadOwnerDuties',
            duties_dot_duties__pb2.ReadOwnerDutiesRequest.SerializeToString,
            duties_dot_duties__pb2.ReadDutiesResponse.FromString,
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
    def ReadTargetDuties(request,
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
            '/duties.Duties/ReadTargetDuties',
            duties_dot_duties__pb2.ReadTargetDutiesRequest.SerializeToString,
            duties_dot_duties__pb2.ReadDutiesResponse.FromString,
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
    def CreateDuty(request,
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
            '/duties.Duties/CreateDuty',
            duties_dot_duties__pb2.CreateDutyRequest.SerializeToString,
            duties_dot_duties__pb2.CreateDutyResponse.FromString,
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
    def UpdateDuty(request,
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
            '/duties.Duties/UpdateDuty',
            duties_dot_duties__pb2.UpdateDutyRequest.SerializeToString,
            duties_dot_duties__pb2.UpdateDutyResponse.FromString,
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
    def DeleteDuty(request,
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
            '/duties.Duties/DeleteDuty',
            duties_dot_duties__pb2.DeleteDutyRequest.SerializeToString,
            duties_dot_duties__pb2.DeleteDutyResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
