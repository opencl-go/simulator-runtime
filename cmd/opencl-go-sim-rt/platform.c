#include "icd.h"

extern cl_int goGetPlatformIDs(cl_uint num_entries, cl_platform_id *platforms, cl_uint *num_platforms);

CL_API_ENTRY cl_int CL_API_CALL
    clGetPlatformIDs(cl_uint num_entries, cl_platform_id *platforms, cl_uint *num_platforms)
{
    return goGetPlatformIDs(num_entries, platforms, num_platforms);
}

extern cl_int goGetPlatformInfo(uint64_t platformHandle, cl_platform_info param_name,
    size_t param_value_size, void *param_value, size_t *param_value_size_ret);

CL_API_ENTRY cl_int CL_API_CALL
    clGetPlatformInfo(cl_platform_id platform, cl_platform_info param_name,
        size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return goGetPlatformInfo(((goDispatchObject*)(platform))->handle, param_name, param_value_size, param_value, param_value_size_ret);
}

CL_API_ENTRY cl_int CL_API_CALL clUnloadPlatformCompiler(cl_platform_id platform)
{
    return CL_OUT_OF_RESOURCES;
}

extern uintptr_t goGetExtensionFunctionAddressForPlatform(uint64_t platformHandle, char *func_name);

CL_API_ENTRY void *CL_API_CALL clGetExtensionFunctionAddressForPlatform(cl_platform_id platform, const char *func_name)
{
    return (void *)(goGetExtensionFunctionAddressForPlatform(((goDispatchObject*)(platform))->handle, (char *)(func_name)));
}
