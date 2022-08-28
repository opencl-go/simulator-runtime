#include "icd.h"

CL_API_ENTRY cl_int CL_API_CALL clRetainMemObject(cl_mem memobj)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clReleaseMemObject(cl_mem memobj)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetMemObjectInfo(cl_mem memobj,
    cl_mem_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueUnmapMemObject(cl_command_queue command_queue,
    cl_mem memobj, void *mapped_ptr,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueMigrateMemObjects(cl_command_queue command_queue,
    cl_uint num_mem_objects, const cl_mem *mem_objects, cl_mem_migration_flags flags,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clSetMemObjectDestructorCallback(cl_mem memobj,
    void (CL_CALLBACK *pfn_notify)(cl_mem memobj, void *user_data), void *user_data)
{
    return CL_OUT_OF_RESOURCES;
}

