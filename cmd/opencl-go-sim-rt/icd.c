#include <memory.h>
#include <string.h>

#include "icd.h"

static cl_icd_dispatch globalDispatch = (struct _cl_icd_dispatch){
    clGetPlatformIDs,
    clGetPlatformInfo,
    clGetDeviceIDs,
    clGetDeviceInfo,
    clCreateContext,
    clCreateContextFromType,
    clRetainContext,
    clReleaseContext,
    clGetContextInfo,
};

goDispatchObject *newDispatchObject(uint64_t handle)
{
    goDispatchObject *obj = (goDispatchObject *)(calloc(1, sizeof(goDispatchObject)));
    obj->dispatch = &globalDispatch;
    obj->handle = handle;
    return obj;
}

CL_API_ENTRY cl_int CL_API_CALL
    clIcdGetPlatformIDsKHR(cl_uint num_entries, cl_platform_id *platforms, cl_uint *num_platforms)
{
    return clGetPlatformIDs(num_entries, platforms, num_platforms);
}

CL_API_ENTRY void *CL_API_CALL clGetExtensionFunctionAddress(const char *func_name)
{
    if (strcmp(func_name, "clIcdGetPlatformIDsKHR") == 0)
    {
        return clIcdGetPlatformIDsKHR;
    }
    return NULL;
}
