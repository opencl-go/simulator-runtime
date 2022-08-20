#include <memory.h>
#include <string.h>

#include "icd.h"
#include "platform.h"

static cl_icd_dispatch globalDispatch = (struct _cl_icd_dispatch){
    clGetPlatformIDs,
    clGetPlatformInfo,
};
static struct _cl_platform_id simulatorPlatform = (struct _cl_platform_id){&globalDispatch, 0};
static struct _cl_platform_id *availablePlatforms[1] =
{
    &simulatorPlatform
};

extern uintptr_t simGetPlatform();

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
        simulatorPlatform.goHandle = simGetPlatform();

        size_t reportedPlatformCount = (availablePlatformCount > num_entries) ? num_entries : availablePlatformCount;
        memset(platforms, 0x00, sizeof(cl_platform_id) * (size_t)(num_entries));
        for (size_t i = 0; i < reportedPlatformCount; i++)
        {
            platforms[i] = availablePlatforms[i];
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

void *CL_API_CALL clGetExtensionFunctionAddress(const char *func_name)
{
    if (strcmp(func_name, "clIcdGetPlatformIDsKHR") == 0)
    {
        return clIcdGetPlatformIDsKHR;
    }
    return NULL;
}
