import requests

resp = requests.post('https://textbelt.com/text', {
  'phone': '9085252880',
  'message': 'Hello world',
  'key': '01bdcca1e50296e27c698e6f673cd7087db326bbH7Qqf9FgSH1acYSd2v6UdEfnc',
})
print(resp.json())