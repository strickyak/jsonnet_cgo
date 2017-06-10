#ifndef LIBJSONNET_BRIDGE_H
#define LIBJSONNET_BRIDGE_H
#include <libjsonnet.h>

typedef JsonnetImportCallback* JsonnetImportCallbackPtr;

typedef JsonnetNativeCallback* JsonnetNativeCallbackPtr;

typedef struct JsonnetJsonValue* JsonnetJsonValuePtr;

struct JsonnetVm* go_get_guts(void* ctx);

char* CallImport_cgo(void *ctx, const char *base, const char *rel, char **found_here, int *success);

JsonnetJsonValuePtr CallNative_cgo(void *ctx, const JsonnetJsonValuePtr const *argv, int *succes);

char* go_call_import(void* vm, char *base, char *rel, char **path, int *success);

char* go_call_native(void* native_context, char *argv, int *success);

// The following are helpers for converting a Go slice of strings
// into an array of null terminated strings.
char** makeCharArray(int size);

void setArrayString(char **a, char *s, int n);

void freeCharArray(char **a, int size);

#endif
