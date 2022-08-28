#include "icd.h"

CL_API_ENTRY cl_mem CL_API_CALL clCreateBuffer(cl_context context,
    cl_mem_flags flags, size_t size, void *host_ptr, cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_mem CL_API_CALL clCreateSubBuffer(cl_mem buffer,
    cl_mem_flags flags, cl_buffer_create_type buffer_create_type, const void *buffer_create_info,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueReadBuffer(cl_command_queue command_queue,
    cl_mem buffer, cl_bool blocking_read,
    size_t offset, size_t size, void *ptr,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueReadBufferRect(cl_command_queue command_queue,
    cl_mem buffer, cl_bool blocking_read,
    const size_t *buffer_origin, const size_t *host_origin, const size_t *region,
    size_t buffer_row_pitch, size_t buffer_slice_pitch, size_t host_row_pitch, size_t host_slice_pitch, void *ptr,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueWriteBuffer(cl_command_queue command_queue,
    cl_mem buffer, cl_bool blocking_write,
    size_t offset, size_t size, const void *ptr,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueWriteBufferRect(cl_command_queue command_queue,
    cl_mem buffer, cl_bool blocking_write,
    const size_t *buffer_origin, const size_t *host_origin, const size_t *region,
    size_t buffer_row_pitch, size_t buffer_slice_pitch, size_t host_row_pitch, size_t host_slice_pitch, const void *ptr,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueFillBuffer(cl_command_queue command_queue,
    cl_mem buffer,
    const void *pattern, size_t pattern_size, size_t offset, size_t size,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueCopyBuffer(cl_command_queue command_queue,
    cl_mem src_buffer, cl_mem dst_buffer,
    size_t src_offset, size_t dst_offset, size_t size,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueCopyBufferRect(cl_command_queue command_queue,
    cl_mem src_buffer, cl_mem dst_buffer,
    const size_t *src_origin, const size_t *dst_origin, const size_t *region,
    size_t src_row_pitch, size_t src_slice_pitch, size_t dst_row_pitch, size_t dst_slice_pitch,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY void *CL_API_CALL clEnqueueMapBuffer(cl_command_queue command_queue,
    cl_mem buffer, cl_bool blocking_map, cl_map_flags map_flags,
    size_t offset, size_t size,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event,
    cl_int *errcode_ret)
{
    return NULL;
}
