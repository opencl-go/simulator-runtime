#include "icd.h"

CL_API_ENTRY cl_mem CL_API_CALL clCreatePipe(cl_context context,
    cl_mem_flags flags, cl_uint pipe_packet_size, cl_uint pipe_max_packets, const cl_pipe_properties *properties,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clGetPipeInfo(cl_mem pipe,
    cl_pipe_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}
