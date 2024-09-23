export const conf = {
  defaultBaseURL: 'http://localhost:8080'
}

const getBaseURL = () => {
  return process.env.SERVICE_URL || conf.defaultBaseURL
}

export const fullURL = (path) => {
  return `${getBaseURL()}${path}`
}
