#include "icd.h"

extern cl_int goGetDeviceIDs(uint64_t platformHandle, cl_device_type device_type,
    cl_uint num_entries, cl_device_id *devices, cl_uint *num_devices);

CL_API_ENTRY cl_int CL_API_CALL
    clGetDeviceIDs(cl_platform_id platform, cl_device_type device_type,
        cl_uint num_entries, cl_device_id *devices, cl_uint *num_devices)
{
    return goGetDeviceIDs(((goDispatchObject*)(platform))->handle, device_type, num_entries, devices, num_devices);
}

extern cl_int goGetDeviceInfo(uint64_t deviceHandle,
    cl_device_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret);

CL_API_ENTRY cl_int CL_API_CALL
    clGetDeviceInfo(cl_device_id device,
        cl_device_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return goGetDeviceInfo(((goDispatchObject*)(device))->handle,
        param_name, param_value_size, param_value, param_value_size_ret);
}

CL_API_ENTRY cl_int CL_API_CALL clCreateSubDevices(cl_device_id in_device,
    const cl_device_partition_property *properties,
    cl_uint num_devices, cl_device_id *out_devices, cl_uint *num_devices_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clRetainDevice(cl_device_id device)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clReleaseDevice(cl_device_id device)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clSetDefaultDeviceCommandQueue(cl_context context,
    cl_device_id device, cl_command_queue command_queue)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetDeviceAndHostTimer(cl_device_id device,
    cl_ulong *device_timestamp, cl_ulong *host_timestamp)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetHostTimer(cl_device_id device,
    cl_ulong *host_timestamp)
{
    return CL_OUT_OF_RESOURCES;
}
