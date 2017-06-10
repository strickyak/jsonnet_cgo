#include <memory.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <libjsonnet.h>
#include "bridge.h"

char* CallImport_cgo(void *ctx, const char *base, const char *rel, char **found_here, int *success) {
  struct JsonnetVm* vm = ctx;
  char *path = NULL;
  char* result = go_call_import(vm, (char*)base, (char*)rel, &path, success);
  if (*success) {
    char* found_here_buf = jsonnet_realloc(vm, NULL, strlen(path)+1);
    strcpy(found_here_buf, path);
    *found_here = found_here_buf;
  }
  char* buf = jsonnet_realloc(vm, NULL, strlen(result)+1);
  strcpy(buf, result);
  return buf;
}

struct NativeContext {
    struct JsonnetVm *vm;
    void *callback_function;
    int arg_count;
};

// This function is bound for every native callback, but with a different context.
JsonnetJsonValuePtr CallNative_cgo(void *ctx, const JsonnetJsonValuePtr const *argv, int *success) {
  struct NativeContext* context = ctx;

  const char** params = calloc(context->arg_count, sizeof(char *));
  int i;
  for (i = 0; i < context->arg_count; i++) {
    params[i] = jsonnet_json_extract_string(context->vm, argv[i]);
  }

  char* result = go_call_native(ctx, (char **)params, success);

  // We free only the array, not the underlying strings which are handled
  // by the jsonnet vm.
  free(params);

  // Currently the return value can be a plain string value.
  return jsonnet_json_make_string(context->vm, result);
}

// The following are helpers for converting a Go slice of strings
// into an array of C strings.
char** makeCharArray(int size) {
  return calloc(size, sizeof(char*));
}

void setArrayString(char **a, char *s, int n) {
  a[n] = s;
}

void freeCharArray(char **a, int size) {
  int i;
  for (i = 0; i < size; i++) {
    free(a[i]);
  }
  free(a);
}
