import { expect } from 'chai';
import { Readable } from 'stream';

describe.only('stream.Readable', () => {
  it('should emit data events with correct data', done => {
    const dataChunks = ['chunk1', 'chunk2', 'chunk3'];
    const expectedData = dataChunks.join('');

    const readableStream = new Readable({
      read() {
        dataChunks.forEach(chunk => {
          this.push(chunk);
        });
        this.push(null); // End of stream
      },
    });

    let receivedData = '';

    readableStream.on('data', chunk => {
      receivedData += chunk.toString();
    });

    readableStream.on('end', () => {
      try {
        expect(receivedData).to.equal(expectedData);
        done();
      } catch (err) {
        done(err);
      }
    });

    readableStream.resume();
  });

  it('should handle highWaterMark option correctly', done => {
    const dataChunks = ['chunk1', 'chunk2', 'chunk3'];
    const highWaterMark = 5; // Set a low highWaterMark value

    const readableStream = new Readable({
      read() {
        dataChunks.forEach(chunk => {
          this.push(chunk);
        });
        this.push(null); // End of stream
      },
      highWaterMark,
    });

    let receivedData = '';

    readableStream.on('data', chunk => {
      receivedData += chunk.toString();
    });

    readableStream.on('end', () => {
      // The data should still be correctly received even with a low highWaterMark
      expect(receivedData).to.equal(dataChunks.join(''));
      done();
    });

    readableStream.resume();
  });

  it('should emit readable event when data can be read', done => {
    const readableStream = new Readable({
      read() {
        // Do not push data immediately
      },
    });

    readableStream.on('readable', () => {
      const chunk = readableStream.read();
      expect(chunk.toString()).to.equal('test');
      done();
    });

    readableStream.push('test');
    readableStream.push(null);
  });

  it('should handle pause and resume correctly', done => {
    const dataChunks = ['chunk1', 'chunk2', 'chunk3'];

    const readableStream = new Readable({
      read() {
        dataChunks.forEach(chunk => {
          this.push(chunk);
        });
        this.push(null); // End of stream
      },
    });

    let receivedData = '';

    readableStream.on('data', chunk => {
      receivedData += chunk.toString();
      readableStream.pause();
      setTimeout(() => {
        readableStream.resume();
      }, 10);
    });

    readableStream.on('end', () => {
      try {
        expect(receivedData).to.equal(dataChunks.join(''));
        done();
      } catch (err) {
        done(err);
      }
    });

    readableStream.resume();
  });

  it('should emit error event for custom error handling', done => {
    const customError = new Error('Custom Error');

    const readableStream = new Readable({
      read() {
        this.emit('error', customError);
      },
    });

    readableStream.on('error', error => {
      expect(error).to.equal(customError);
      done();
    });
  });

  it('should handle empty stream', done => {
    const readableStream = new Readable({
      read() {
        // Do not push any data
        this.push(null);
      },
    });

    readableStream.on('readable', () => {
      const chunk = readableStream.read();
      expect(chunk).to.be.null;
      done();
    });
  });

  it('should handle simultaneous read and push', done => {
    const readableStream = new Readable({
      read() {
        this.push('data1');
        this.push('data2');
        this.push(null); // End of stream
      },
    });

    let receivedData = '';

    readableStream.on('readable', () => {
      const chunk = readableStream.read();
      if (chunk !== null) {
        receivedData += chunk.toString();
      }
    });

    readableStream.on('end', () => {
      expect(receivedData).to.equal('data1data2');
      done();
    });
  });

  it('should handle backpressure', done => {
    const dataChunks = ['chunk1', 'chunk2', 'chunk3'];
    const highWaterMark = 5;

    const readableStream = new Readable({
      read() {
        dataChunks.forEach(chunk => {
          this.push(chunk);
        });
        this.push(null); // End of stream
      },
      highWaterMark,
    });

    let receivedData = '';

    readableStream.on('readable', () => {
      const chunk = readableStream.read();
      if (chunk !== null) {
        receivedData += chunk.toString();
      }
    });

    readableStream.on('end', () => {
      expect(receivedData).to.equal(dataChunks.join(''));
      done();
    });
  });

  it('should handle stream destruction', done => {
    const readableStream = new Readable({
      read() {
        this.push('data');
        this.push(null); // End of stream
      },
    });

    readableStream.on('data', () => {
      readableStream.destroy();
    });

    readableStream.on('close', done);

    readableStream.resume();
  });

  it('should emit error event when source has an error', done => {
    const readableStream = new Readable({
      read() {
        this.emit('error', new Error('Source Error'));
      },
    });

    readableStream.on('error', error => {
      expect(error.message).to.equal('Source Error');
      done();
    });

    readableStream.resume();
  });

  it('should emit error event during stream processing', done => {
    const readableStream = new Readable({
      read() {
        this.emit('error', new Error('Stream Processing Error'));
      },
    });

    readableStream.on('error', error => {
      expect(error.message).to.equal('Stream Processing Error');
      done();
    });

    readableStream.resume();
  });

  it('should handle multiple consumers', done => {
    const readableStream = new Readable({
      read() {
        this.push('data');
        this.push(null); // End of stream
      },
    });

    let consumer1Data = '';
    let consumer2Data = '';

    readableStream.on('data', chunk => {
      consumer1Data += chunk.toString();
    });

    readableStream.on('data', chunk => {
      consumer2Data += chunk.toString();
    });

    readableStream.on('end', () => {
      expect(consumer1Data).to.equal('data');
      expect(consumer2Data).to.equal('data');
      done();
    });

    readableStream.resume();
  });

  it('should handle object mode', done => {
    const objectStream = new Readable({
      read() {
        this.push({ key: 'value' });
        this.push(null); // End of stream
      },
      objectMode: true,
    });

    let receivedData = '';

    objectStream.on('data', data => {
      receivedData = data;
    });

    objectStream.on('end', () => {
      expect(receivedData).to.deep.equal({ key: 'value' });
      done();
    });

    objectStream.resume();
  });

  it('should handle large amounts of data without memory issues', function (done) {
    // Generate a large amount of data (e.g., 100 MB)
    const largeDataSize = 1024 * 1024 * 100;
    const largeData = Buffer.alloc(largeDataSize);

    // Counter to track the number of data events
    let dataEventCount = 0;
    let i = 0;
    const readableStream = new Readable({
      read() {
        // Push the large data in chunks
        for (; i < largeData.length; i += 1024) {
          if (!this.push(largeData.subarray(i, i + 1024))) {
            i += 1024;
            return;
          }
        }

        // Simulate an asynchronous process (e.g., reading from a file)
        setTimeout(() => {
          // Signal the end of the stream
          this.push(null);
        }, 100);
      },
    });

    readableStream.on('data', chunk => {
      // Count the number of data events
      dataEventCount++;
    });

    readableStream.on('end', () => {
      // Check if the number of data events is reasonable for the large data size
      // Adjust this threshold based on your specific scenario
      const reasonableDataEventCount = largeDataSize / 1024;

      expect(dataEventCount).to.be.at.least(reasonableDataEventCount);

      done();
    });

    readableStream.resume();
  });

  it('should emit custom event when a specific condition is met', done => {
    const readableStream = new Readable({
      read() {
        // Stream logic
      },
    });

    readableStream.on('customEvent', data => {
      // Custom event handling logic
      expect(data).to.equal('someData');
      done();
    });

    // Trigger custom event in the stream logic
    readableStream.emit('customEvent', 'someData');

    readableStream.resume();
  });
});
