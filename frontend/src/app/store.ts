import { configureStore } from '@reduxjs/toolkit'

// Reducers (to be created)
const appSlice = {
  name: 'app',
  initialState: {
    isLoading: false,
    error: null,
  },
  reducers: {},
}

export const store = configureStore({
  reducer: {
    app: () => appSlice.initialState, // Temporary reducer
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware({
      serializableCheck: {
        ignoredActions: [],
      },
    }),
  devTools: process.env.NODE_ENV !== 'production',
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch