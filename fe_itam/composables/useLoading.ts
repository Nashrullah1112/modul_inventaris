const isVisible = ref(false)

export function useLoading() {
  const showLoading = (): void => {
    isVisible.value = true
  }

  const hideLoading = (): void => {
    isVisible.value = false
  }

  return {
    isVisible,
    showLoading,
    hideLoading
  }
}