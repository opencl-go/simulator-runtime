#pragma once

#if defined(_WIN32)
#define CL_API_ENTRY __declspec(dllexport)
#endif
#include "api.h"

typedef struct {
    cl_icd_dispatch *dispatch;
    uint64_t handle;
} goDispatchObject;

extern goDispatchObject *newDispatchObject(uint64_t handle);
