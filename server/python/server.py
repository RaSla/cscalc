import flask
# from redis import Redis
import datetime
import os
import socket

app = flask.Flask(__name__)
host = socket.gethostname()

# Connect to Redis
#redis_pass = os.getenv('REDIS_PASSWORD', '')
#redis = Redis(host='redis', port=6379, db=0, password=redis_pass, socket_timeout=None, connection_pool=None, charset='utf-8')
# redis = Redis(host='redis', port=6379, db=0)

@app.route('/')
def hi():
    dt_iso = datetime.datetime.utcnow().replace(microsecond=0, tzinfo=datetime.timezone.utc).isoformat()
    txt = 'Hello World from Python & Flask!\n DateTime (UTC): "%s"\n My hostname is "%s"\n\n/py3/hits - try Redis hits-counter\n' % (dt_iso, host)
    return flask.Response(txt, mimetype='text/plain')

@app.route('/plus')
def plus():
    a = flask.request.args.get('a')
    b = flask.request.args.get('b')
    res = int(a) + int(b)
    txt = 'a = ' + a + ', b = ' + b + '\na + b = ' + str(res) + '\n'
    return flask.Response(txt, mimetype='text/plain')

@app.route('/minus')
def minus():
    a = flask.request.args.get('a')
    b = flask.request.args.get('b')
    res = int(a) - int(b)
    txt = 'a = ' + a + ', b = ' + b + '\na - b = ' + str(res) + '\n'
    return flask.Response(txt, mimetype='text/plain')

@app.route('/api/minus')
def api_minus():
    a = flask.request.args.get('a')
    b = flask.request.args.get('b')
    res = int(a) - int(b)
    txt = '{"result": "' + str(res) + '"}\n'
    return flask.Response(txt, mimetype='text/json')

@app.route('/api/plus')
def api_plus():
    a = flask.request.args.get('a')
    b = flask.request.args.get('b')
    res = int(a) + int(b)
    txt = '{"result": "' + str(res) + '"}\n'
    return flask.Response(txt, mimetype='text/json')

@app.route('/multiply')
def multiply():
    a = flask.request.args.get('a')
    b = flask.request.args.get('b')
    res = int(a) * int(b)
    txt = 'a = ' + a + ', b = ' + b + '\na * b = ' + str(res) + '\n'
    return flask.Response(txt, mimetype='text/plain')

@app.route('/api/multiply')
def api_multiply():
    a = flask.request.args.get('a')
    b = flask.request.args.get('b')
    res = int(a) * int(b)
    txt = '{"result": "' + str(res) + '"}\n'
    return flask.Response(txt, mimetype='text/json')

@app.route('/divide')
def divide():
    a = flask.request.args.get('a')
    b = flask.request.args.get('b')
    res = int(a) / int(b)
    txt = 'a = ' + a + ', b = ' + b + '\na / b = ' + str(res) + '\n'
    return flask.Response(txt, mimetype='text/plain')

@app.route('/api/divide')
def api_divide():
    a = flask.request.args.get('a')
    b = flask.request.args.get('b')
    res = int(a) / int(b)
    txt = '{"result": "' + str(res) + '"}\n'
    return flask.Response(txt, mimetype='text/json')
# @app.route('/py3/hits')
# def hits():
#     dt_iso = datetime.datetime.utcnow().replace(microsecond=0, tzinfo=datetime.timezone.utc).isoformat()
#     hits = redis.incr('hits')
#     txt = 'Python says:\n DateTime (UTC): "%s"\n My hostname is "%s"\n\n Redis "hits": %s\n' % (dt_iso, host, hits)
#     return Response(txt, mimetype='text/plain')

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=8080, debug=False)
