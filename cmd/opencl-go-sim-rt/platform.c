#include "platform.h"

extern cl_int goGetPlatformInfo(uintptr_t platformHandle, cl_platform_info param_name,
    size_t param_value_size, void *param_value, size_t *param_value_size_ret);

CL_API_ENTRY cl_int CL_API_CALL
    clGetPlatformInfo(cl_platform_id platform, cl_platform_info param_name,
        size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return goGetPlatformInfo(platform->goHandle, param_name, param_value_size, param_value, param_value_size_ret);
}

extern cl_int goGetDeviceIDs(uintptr_t platformHandle, cl_device_type device_type,
    cl_uint num_entries, cl_device_id *devices, cl_uint *num_devices);

CL_API_ENTRY cl_int CL_API_CALL
    clGetDeviceIDs(cl_platform_id platform, cl_device_type device_type,
        cl_uint num_entries, cl_device_id *devices, cl_uint *num_devices)
{
    return goGetDeviceIDs(platform->goHandle, device_type, num_entries, devices, num_devices);
}
