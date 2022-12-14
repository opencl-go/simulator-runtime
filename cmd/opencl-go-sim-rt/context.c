#include "icd.h"

CL_API_ENTRY cl_context CL_API_CALL clCreateContext(const cl_context_properties *properties,
    cl_uint num_devices, const cl_device_id *devices,
    void (CL_CALLBACK *pfn_notify)(const char *errinfo, const void *private_info, size_t cb, void *user_data),
    void *user_data,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_context CL_API_CALL clCreateContextFromType(const cl_context_properties *properties,
    cl_device_type device_type,
    void (CL_CALLBACK *pfn_notify)(const char *errinfo, const void *private_info, size_t cb, void *user_data),
    void *user_data,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clRetainContext(cl_context context)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clReleaseContext(cl_context context)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetContextInfo(cl_context context,
    cl_context_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clSetContextDestructorCallback(cl_context context,
    void (CL_CALLBACK *pfn_notify)(cl_context context, void *user_data), void *user_data)
{
    return CL_OUT_OF_RESOURCES;
}
