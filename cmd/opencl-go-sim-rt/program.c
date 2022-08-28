#include "icd.h"

CL_API_ENTRY cl_program CL_API_CALL clCreateProgramWithSource(cl_context context,
    cl_uint count, const char **strings, const size_t *lengths,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_program CL_API_CALL clCreateProgramWithBinary(cl_context context,
    cl_uint num_devices, const cl_device_id *device_list, const size_t *lengths,
    const unsigned char **binaries, cl_int *binary_status,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clRetainProgram(cl_program program)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clReleaseProgram(cl_program program)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clBuildProgram(cl_program program,
    cl_uint num_devices, const cl_device_id * device_list, const char *options,
    void (CL_CALLBACK *pfn_notify)(cl_program program, void *user_data), void *user_data)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clCompileProgram(cl_program program,
    cl_uint num_devices, const cl_device_id *device_list,
    const char *options,
    cl_uint num_input_headers, const cl_program *input_headers, const char **header_include_names,
    void (CL_CALLBACK *pfn_notify)(cl_program program, void *user_data), void *user_data)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_program CL_API_CALL clLinkProgram(cl_context context,
    cl_uint num_devices, const cl_device_id *device_list,
    const char *options,
    cl_uint num_input_programs, const cl_program *input_programs,
    void (CL_CALLBACK *pfn_notify)(cl_program program, void *user_data), void *user_data,
    cl_int *errcode_ret)
{
    return NULL;
}

CL_API_ENTRY cl_int CL_API_CALL clGetProgramInfo(cl_program program,
    cl_program_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clGetProgramBuildInfo(cl_program program, cl_device_id device,
    cl_program_build_info param_name, size_t param_value_size, void *param_value, size_t *param_value_size_ret)
{
    return CL_OUT_OF_RESOURCES;
}
