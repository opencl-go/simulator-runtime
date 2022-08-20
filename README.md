# Simulator runtime

This is an OpenCL runtime implementation that adheres to the ICD loader specification as per
[cl_khr_icd](https://registry.khronos.org/OpenCL/sdk/3.0/docs/man/html/cl_khr_icd.html).

The intention is to provide a simulator that can be used during testing of the wrapper libraries.

**This is work in progress**

## Build and install

You need to build the runtime library and register it as a potential vendor runtime.

### Prerequisites
To build  this library, you need an OpenCL SDK installed on your system.
Refer to [the documentation on opencl-go][opencl-go] on how to do this.

[opencl-go]: https://opencl-go.github.com

### Linux

```
cd cmd/opencl-go-sim-rt
go build -buildmode c-shared .
mkdir --parents /etc/OpenCL/vendors
echo "(path-to-so)/opencl-go-sim-rt" > /etc/OpenCL/vendors/opencl-go-simulator.icd
```

## References

* `cl_khr_icd` extension specification: [https://registry.khronos.org/OpenCL/sdk/3.0/docs/man/html/cl_khr_icd.html](https://registry.khronos.org/OpenCL/sdk/3.0/docs/man/html/cl_khr_icd.html)
* Khronos ICD Loader: [https://github.com/KhronosGroup/OpenCL-ICD-Loader](https://github.com/KhronosGroup/OpenCL-ICD-Loader)

## License

This project is based on the MIT License. See `LICENSE` file.
