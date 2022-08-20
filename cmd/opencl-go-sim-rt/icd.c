#include <memory.h>
#include <string.h>

#if defined(_WIN32)
#define CL_API_ENTRY __declspec(dllexport)
#endif
#include "api.h"

struct _cl_platform_id
{
    cl_icd_dispatch *dispatch;
};

static cl_icd_dispatch globalDispatch = (struct _cl_icd_dispatch){
    clGetPlatformIDs,
    clGetPlatformInfo,
};
static struct _cl_platform_id availablePlatforms[1] =
{
    (struct _cl_platform_id){&globalDispatch}
};

CL_API_ENTRY cl_int CL_API_CALL
    clGetPlatformIDs(cl_uint num_entries, cl_platform_id *platforms, cl_uint *num_platforms)
{
    if (((num_entries > 0) && (platforms == NULL)) || ((platforms == NULL) && (num_platforms == NULL)))
    {
        return CL_INVALID_VALUE;
    }
    size_t availablePlatformCount = (sizeof(availablePlatforms) / sizeof(availablePlatforms[0]));
    if (num_entries > 0)
    {
        size_t reportedPlatformCount = (availablePlatformCount > num_entries) ? num_entries : availablePlatformCount;
        memset(platforms, 0x00, sizeof(cl_platform_id) * (size_t)(num_entries));
        for (size_t i = 0; i < reportedPlatformCount; i++)
        {
            platforms[i] = &availablePlatforms[i];
        }
    }
    if (num_platforms != NULL)
    {
        *num_platforms = (cl_uint)(availablePlatformCount);
    }
    return CL_SUCCESS;
}

CL_API_ENTRY cl_int CL_API_CALL
    clIcdGetPlatformIDsKHR(cl_uint num_entries, cl_platform_id *platforms, cl_uint *num_platforms)
{
    return clGetPlatformIDs(num_entries, platforms, num_platforms);
}

extern cl_int simGetPlatformInfo(cl_platform_id platform, cl_platform_info param_name,
    size_t param_value_size, void *param_value, size_t *param_value_size_ret);

CL_API_ENTRY cl_int CL_API_CALL
    clGetPlatformInfo(cl_platform_id platform, cl_platform_info param_name,
        size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return simGetPlatformInfo(platform, param_name, param_value_size, param_value, param_value_size_ret);
}

void *CL_API_CALL clGetExtensionFunctionAddress(const char *func_name)
{
    if (strcmp(func_name, "clIcdGetPlatformIDsKHR") == 0)
    {
        return clIcdGetPlatformIDsKHR;
    }
    return NULL;
}
