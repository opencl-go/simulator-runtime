#include "icd.h"

CL_API_ENTRY cl_sampler CL_API_CALL clCreateSampler(cl_context context,
    cl_bool normalized_coords, cl_addressing_mode addressing_mode, cl_filter_mode filter_mode,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_sampler CL_API_CALL clCreateSamplerWithProperties(cl_context context,
    const cl_sampler_properties *sampler_properties,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clRetainSampler(cl_sampler sampler)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clReleaseSampler(cl_sampler sampler)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetSamplerInfo(cl_sampler sampler,
    cl_sampler_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}
