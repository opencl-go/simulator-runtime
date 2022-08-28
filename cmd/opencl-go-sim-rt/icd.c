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
    clCreateCommandQueue,
    clRetainCommandQueue,
    clReleaseCommandQueue,
    clGetCommandQueueInfo,
    NULL, // clSetCommandQueueProperty
    clCreateBuffer,
    NULL, // clCreateImage2D
    NULL, // clCreateImage3D
    clRetainMemObject,
    clReleaseMemObject,
    clGetSupportedImageFormats,
    clGetMemObjectInfo,
    clGetImageInfo,
    clCreateSampler,
    clRetainSampler,
    clReleaseSampler,
    clGetSamplerInfo,
    clCreateProgramWithSource,
    clCreateProgramWithBinary,
    clRetainProgram,
    clReleaseProgram,
    clBuildProgram,
    NULL, // clUnloadCompiler
    clGetProgramInfo,
    clGetProgramBuildInfo,
    clCreateKernel,
    clCreateKernelsInProgram,
    clRetainKernel,
    clReleaseKernel,
    clSetKernelArg,
    clGetKernelInfo,
    clGetKernelWorkGroupInfo,
    clWaitForEvents,
    clGetEventInfo,
    clRetainEvent,
    clReleaseEvent,
    clGetEventProfilingInfo,
    clFlush,
    clFinish,
    clEnqueueReadBuffer,
    clEnqueueWriteBuffer,
    clEnqueueCopyBuffer,
    clEnqueueReadImage,
    clEnqueueWriteImage,
    clEnqueueCopyImage,
    clEnqueueCopyImageToBuffer,
    clEnqueueCopyBufferToImage,
    clEnqueueMapBuffer,
    clEnqueueMapImage,
    clEnqueueUnmapMemObject,
    clEnqueueNDRangeKernel,
    clEnqueueTask,
    clEnqueueNativeKernel,
    NULL, // clEnqueueMarker
    NULL, // clEnqueueWaitForEvents
    NULL, // clEnqueueBarrier
    clGetExtensionFunctionAddress,

    NULL, // clCreateFromGLBuffer
    NULL, // clCreateFromGLTexture2D
    NULL, // clCreateFromGLTexture3D
    NULL, // clCreateFromGLRenderbuffer
    NULL, // clGetGLObjectInfo
    NULL, // clGetGLTextureInfo
    NULL, // clEnqueueAcquireGLObjects
    NULL, // clEnqueueReleaseGLObjects
    NULL, // clGetGLContextInfoKHR

    NULL, // clGetDeviceIDsFromD3D10KHR
    NULL, // clCreateFromD3D10BufferKHR
    NULL, // clCreateFromD3D10Texture2DKHR
    NULL, // clCreateFromD3D10Texture3DKHR
    NULL, // clEnqueueAcquireD3D10ObjectsKHR
    NULL, // clEnqueueReleaseD3D10ObjectsKHR

    clSetEventCallback,
    clCreateSubBuffer,
    clSetMemObjectDestructorCallback,
    clCreateUserEvent,
    clSetUserEventStatus,
    clEnqueueReadBufferRect,
    clEnqueueWriteBufferRect,
    clEnqueueCopyBufferRect,

    NULL, // clCreateSubDevicesEXT
    NULL, // clRetainDeviceEXT
    NULL, // clReleaseDeviceEXT

    NULL, // clCreateEventFromGLsyncKHR

    clCreateSubDevices,
    clRetainDevice,
    clReleaseDevice,
    clCreateImage,
    clCreateProgramWithBuiltInKernels,
    clCompileProgram,
    clLinkProgram,
    clUnloadPlatformCompiler,
    clGetKernelArgInfo,
    clEnqueueFillBuffer,
    clEnqueueFillImage,
    clEnqueueMigrateMemObjects,
    clEnqueueMarkerWithWaitList,
    clEnqueueBarrierWithWaitList,
    clGetExtensionFunctionAddressForPlatform,

    NULL, // clCreateFromGLTexture
    NULL, // clGetDeviceIDsFromD3D11KHR
    NULL, // clCreateFromD3D11Texture2DKHR
    NULL, // clCreateFromD3D11Texture3DKHR
    NULL, // clCreateFromDX9MediaSurfaceKHR
    NULL, // clEnqueueAcquireD3D11ObjectsKHR
    NULL, // clEnqueueReleaseD3D11ObjectsKHR

    NULL, // clGetDeviceIDsFromDX9MediaAdapterKHR
    NULL, // clEnqueueAcquireDX9MediaSurfacesKHR
    NULL, // clEnqueueReleaseDX9MediaSurfacesKHR

    NULL, // clCreateFromEGLImageKHR
    NULL, // clEnqueueAcquireEGLObjectsKHR
    NULL, // clEnqueueReleaseEGLObjectsKHR

    NULL, // clCreateEventFromEGLSyncKHR
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
