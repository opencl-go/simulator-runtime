#include "icd.h"

CL_API_ENTRY cl_int CL_API_CALL clWaitForEvents(cl_uint num_events, const cl_event *event_list)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetEventInfo(cl_event event,
    cl_event_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_event CL_API_CALL clCreateUserEvent(cl_context context, cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clRetainEvent(cl_event event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clReleaseEvent(cl_event event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clSetUserEventStatus(cl_event event, cl_int execution_status)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clSetEventCallback(cl_event event,
    cl_int command_exec_callback_type,
    void (CL_CALLBACK *pfn_notify)(cl_event event, cl_int event_command_status, void *user_data), void *user_data)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetEventProfilingInfo(cl_event event,
    cl_profiling_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueMarkerWithWaitList(cl_command_queue command_queue,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueBarrierWithWaitList(cl_command_queue command_queue,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}
