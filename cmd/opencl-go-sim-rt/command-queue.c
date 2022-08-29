#include "icd.h"

CL_API_ENTRY cl_command_queue CL_API_CALL clCreateCommandQueue(cl_context context,
    cl_device_id device, cl_command_queue_properties properties,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_command_queue CL_API_CALL clCreateCommandQueueWithProperties(cl_context context,
    cl_device_id device, const cl_queue_properties *properties,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clRetainCommandQueue(cl_command_queue command_queue)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clReleaseCommandQueue(cl_command_queue command_queue)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetCommandQueueInfo(cl_command_queue command_queue,
    cl_command_queue_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clFlush(cl_command_queue command_queue)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clFinish(cl_command_queue command_queue)
{
    return CL_OUT_OF_RESOURCES;
}
