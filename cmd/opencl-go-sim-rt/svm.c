#include "icd.h"

CL_API_ENTRY void *CL_API_CALL clSVMAlloc(cl_context context,
    cl_svm_mem_flags flags, size_t size, cl_uint alignment)
{
    return NULL;
}

CL_API_ENTRY void CL_API_CALL clSVMFree(cl_context context,
    void *svm_pointer)
{
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueSVMFree(cl_command_queue command_queue,
    cl_uint num_svm_pointers, void *svm_pointers[],
    void (CL_CALLBACK * pfn_free_func)(cl_command_queue queue,
        cl_uint num_svm_pointers, void *svm_pointers[], void *user_data), void *user_data,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueSVMMemcpy(cl_command_queue command_queue,
    cl_bool blocking_copy,
    void *dst_ptr, const void *src_ptr, size_t size,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueSVMMemFill(cl_command_queue command_queue,
    void *svm_ptr, const void *pattern, size_t pattern_size, size_t size,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueSVMMap(cl_command_queue command_queue,
    cl_bool blocking_map, cl_map_flags flags, void *svm_ptr, size_t size,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueSVMUnmap(cl_command_queue command_queue,
    void *svm_ptr,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}

CL_API_ENTRY cl_int CL_API_CALL clEnqueueSVMMigrateMem(cl_command_queue command_queue,
    cl_uint num_svm_pointers, const void **svm_pointers, const size_t *sizes, cl_mem_migration_flags flags,
    cl_uint num_events_in_wait_list, const cl_event *event_wait_list, cl_event *event)
{
    return CL_OUT_OF_RESOURCES;
}
