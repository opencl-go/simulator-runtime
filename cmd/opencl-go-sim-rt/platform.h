#pragma once

#include "icd.h"

struct _cl_platform_id
{
    cl_icd_dispatch *dispatch;
    uintptr_t goHandle;
};
