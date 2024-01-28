import { Writable } from 'stream';
import { expect } from 'chai';

describe.only('stream.Writable', () => {
  it('should receive data correctly', done => {
    const chunks: Buffer[] = [];

    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        chunks.push(chunk);
        callback();
      },
      final(callback) {
        // Test that the received data is correct
        expect(Buffer.concat(chunks).toString()).to.equal('Hello, world!');
        callback();
      },
    });

    // Write data to the writable stream
    writableStream.write('Hello, ');
    writableStream.write('world!');
    writableStream.end(() => {
      done();
    });
  });

  it('should handle backpressure', done => {
    const writableStream = new Writable({
      highWaterMark: 10, // Set a low highWaterMark value
      write(chunk, encoding, callback) {
        // Simulate a slow consumer by delaying the callback
        setTimeout(() => {
          callback();
        }, 10);
      },
    });

    // Write a large amount of data to trigger backpressure
    const largeData = 'x'.repeat(100);
    let i = 0;
    const next = () => {
      for (; i < 10; i++) {
        if (!writableStream.write(largeData)) {
          i++;
          return;
        }
      }
    };

    writableStream.on('drain', () => {
      // The stream is ready to receive more data
      writableStream.end();
      done();
    });

    next();
  });

  it('should handle the finish event', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Writing logic
        callback();
      },
    });

    // Write data to the writable stream
    writableStream.write('Hello, ');
    writableStream.write('world!');

    // Listen for the finish event
    writableStream.on('finish', () => {
      done();
    });

    // End the writable stream to trigger the finish event
    writableStream.end();
  });

  it('should handle errors during writing', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Simulate an error during writing
        callback(new Error('Write Error'));
      },
    });

    // Listen for the error event
    writableStream.on('error', error => {
      expect(error.message).to.equal('Write Error');
      done();
    });

    // Write data to the writable stream
    writableStream.write('Hello, world!');
    writableStream.end();
  });

  it('should handle the close event', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Writing logic
        callback();
      },
    });

    // Listen for the close event
    writableStream.on('close', () => {
      done();
    });

    // End the writable stream to trigger the close event
    writableStream.end();
  });

  it('should handle write after end', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Writing logic
        callback();
      },
    });

    writableStream.end();

    writableStream.on('error', err => {
      expect(err?.message).to.equal('write after end');
      done();
    });

    // Attempt to write data after the stream has ended
    writableStream.write('data');
  });

  it('should buffer data correctly based on highWaterMark', done => {
    const writableStream = new Writable({
      highWaterMark: 10,
      write(chunk, encoding, callback) {
        // Writing logic
        setTimeout(() => callback(), 10);
      },
    });

    writableStream.write('data0');

    // Write data exceeding the highWaterMark
    expect(writableStream.write('data1')).to.be.false;

    // Listen for 'drain' event to check if the stream is ready to receive more data
    writableStream.on('drain', () => {
      expect(writableStream.write('data2')).to.be.true;
      done();
    });
  });

  it('should emit custom event when a specific condition is met', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Writing logic
        callback();
      },
    });

    writableStream.on('customEvent', data => {
      // Custom event handling logic
      expect(data).to.equal('someData');
      done();
    });

    // Trigger custom event in the writing logic
    writableStream.emit('customEvent', 'someData');
  });

  it('should handle object mode', done => {
    const writableStream = new Writable({
      objectMode: true,
      write(chunk, encoding, callback) {
        // Writing logic for object mode
        callback();
      },
    });

    // Write object to the writable stream
    writableStream.write({ key: 'value' });
    writableStream.end(() => {
      done();
    });
  });

  it('should handle concurrent writes', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Writing logic
        setTimeout(() => {
          callback();
        }, 100);
      },
    });

    // Perform multiple concurrent writes
    const writePromises = [];
    for (let i = 0; i < 5; i++) {
      writePromises.push(
        new Promise(resolve => {
          writableStream.write(`data${i}`, resolve);
        })
      );
    }

    // Ensure all writes are completed before ending the stream
    Promise.all(writePromises).then(() => {
      writableStream.end(() => {
        done();
      });
    });
  });

  it('should handle errors during end', done => {
    const writableStream = new Writable({
      final(callback) {
        callback(new Error('Write Error'));
      },
    });

    writableStream.on('error', error => {
      expect(error.message).to.equal('Write Error');
      done();
    });

    writableStream.end();
  });

  it('should handle immediate end without writing data', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Writing logic
        callback();
      },
    });

    writableStream.end(() => {
      done();
    });
  });

  it('should handle writing an empty chunk', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Test that writing an empty chunk is allowed
        callback();
      },
    });

    writableStream.write('', 'utf-8', () => {
      done();
    });
    writableStream.end();
  });

  it('should handle errors in write callback', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Simulate an error in the write callback
        callback(new Error('Write Callback Error'));
      },
    });

    writableStream.on('error', error => {
      expect(error.message).to.equal('Write Callback Error');
      done();
    });

    writableStream.write('data');
    writableStream.end();
  });

  it('should handle concurrent writes followed by end', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Writing logic
        callback();
      },
    });

    const writePromises = [];

    // Perform multiple concurrent writes
    for (let i = 0; i < 5; i++) {
      writePromises.push(
        new Promise(resolve => {
          writableStream.write(`data${i}`, resolve);
        })
      );
    }

    // End the stream after all writes are completed
    Promise.all(writePromises).then(() => {
      writableStream.end(() => {
        done();
      });
    });
  });

  it('should handle multiple end calls', done => {
    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Writing logic
        callback();
      },
    });

    writableStream.end(() => {
      // Subsequent end call should not throw an error
      writableStream.end(() => {
        done();
      });
    });
  });

  it('should handle a large amount of data without memory issues', done => {
    const largeDataSize = 1024 * 1024 * 100; // 100 MB
    const largeData = Buffer.alloc(largeDataSize, 'x');

    const writableStream = new Writable({
      write(chunk, encoding, callback) {
        // Simulate an asynchronous write operation
        setTimeout(() => {
          // Writing logic, for example, write to a file or a database
          // In a real-world scenario, you would handle the chunk appropriately

          // For the test, we'll just acknowledge the write
          callback();
        }, 10);
      },
    });

    // Listen for the 'finish' event to know when all data has been written
    writableStream.on('finish', () => {
      done();
    });

    // Write the large amount of data to the stream
    writableStream.write(largeData);

    // End the stream to trigger the 'finish' event
    writableStream.end();
  }).timeout(5000); // Adjust the timeout based on the expected completion time
});
