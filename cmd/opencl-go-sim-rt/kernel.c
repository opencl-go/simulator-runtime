#include "icd.h"

CL_API_ENTRY cl_kernel CL_API_CALL clCreateKernel(cl_program program,
    const char *kernel_name,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clCreateKernelsInProgram(cl_program program,
    cl_uint num_kernels, cl_kernel *kernels,
    cl_uint *num_kernels_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_program CL_API_CALL clCreateProgramWithBuiltInKernels(cl_context context,
    cl_uint num_devices, const cl_device_id *device_list, const char *kernel_names,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clRetainKernel(cl_kernel kernel)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clReleaseKernel(cl_kernel kernel)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clSetKernelArg(cl_kernel kernel,
    cl_uint arg_index, size_t arg_size, const void *arg_value)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetKernelInfo(cl_kernel kernel,
    cl_kernel_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetKernelArgInfo(cl_kernel kernel,
    cl_uint arg_indx, cl_kernel_arg_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetKernelWorkGroupInfo(cl_kernel kernel, cl_device_id device,
    cl_kernel_work_group_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueNDRangeKernel(cl_command_queue command_queue,
    cl_kernel kernel,
    cl_uint work_dim, const size_t *global_work_offset, const size_t *global_work_size, const size_t *local_work_size,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueTask(cl_command_queue command_queue,
    cl_kernel kernel,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueNativeKernel(cl_command_queue command_queue,
    void (CL_CALLBACK *user_func)(void *),
    void *args, size_t cb_args,
    cl_uint num_mem_objects, const cl_mem *mem_list, const void **args_mem_loc,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}
