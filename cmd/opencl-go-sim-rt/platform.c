#include "icd.h"

extern cl_int simGetPlatformInfo(cl_platform_id platform, cl_platform_info param_name,
    size_t param_value_size, void *param_value, size_t *param_value_size_ret);

CL_API_ENTRY cl_int CL_API_CALL
    clGetPlatformInfo(cl_platform_id platform, cl_platform_info param_name,
        size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return simGetPlatformInfo(platform, param_name, param_value_size, param_value, param_value_size_ret);
}
