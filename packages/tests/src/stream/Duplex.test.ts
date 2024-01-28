import { Duplex } from 'stream';
import { EventEmitter } from 'events';
import { expect } from 'chai';

describe.only('stream.Duplex', () => {
  it('should correctly handle data flow through the Duplex stream', done => {
    const duplexStream = new Duplex({
      read(size) {
        // Implement read logic, for example, pushing data to the readable side
        const data = Buffer.from('Hello, world!');
        this.push(data);
        this.push(null); // Signal the end of data
      },
      write(chunk, encoding, callback) {
        // Implement write logic, for example, handling the writable side
        expect(chunk.toString()).to.equal('Received data');
        callback();
      },
    });

    let receivedData = '';

    duplexStream.on('data', chunk => {
      receivedData += chunk.toString();
    });

    duplexStream.on('end', () => {
      try {
        expect(receivedData).to.equal('Hello, world!');
        done();
      } catch (err) {
        done(err);
      }
    });

    duplexStream.write('Received data');
    duplexStream.end();
    duplexStream.resume();
  });

  it('should handle backpressure in both readable and writable sides', done => {
    const duplexStream = new Duplex({
      read(size) {
        // Implement read logic with backpressure handling
        if (this.push('A'.repeat(1024))) {
          // If push returns true, there's more room in the buffer
          this.push(null); // Signal the end of data
        }
      },
      write(chunk, encoding, callback) {
        // Implement write logic with backpressure handling
        setTimeout(() => {
          callback();
        }, 100); // Simulate an asynchronous write operation
      },
    });

    duplexStream.cork(); // Cork the stream to accumulate data in the buffer

    duplexStream.write('Data1');
    duplexStream.write('Data2');
    duplexStream.write('Data3');

    // Uncork the stream to allow data to flow through
    duplexStream.uncork();

    duplexStream.on('data', chunk => {
      // Handle the readable side data
    });

    duplexStream.on('finish', () => {
      // The writable side finished writing
      done();
    });

    duplexStream.end();
  });

  it('should handle errors in both readable and writable sides', done => {
    const duplexStream = new Duplex({
      read(size) {
        // Implement read logic with error handling
        this.emit('error', new Error('Read Error'));
      },
      write(chunk, encoding, callback) {
        // Implement write logic with error handling
        callback(new Error('Write Error'));
      },
    });

    let hasDoneReadError = false;
    duplexStream.on('error', error => {
      // Handle errors from both sides
      if (!hasDoneReadError) {
        expect(error.message).to.equal('Read Error');
        duplexStream.write('Data');
        duplexStream.end();
        hasDoneReadError = true;
      } else {
        expect(error.message).to.equal('Write Error');
        done();
      }
    });

    duplexStream.resume();
  });

  it('should handle multiple write and read operations', done => {
    const events = new EventEmitter();
    let written = false;

    const duplexStream = new Duplex({
      read(size) {
        if (!written) {
          written = true;
          let i = 0;
          events.on('data', (chunk: any) => {
            i++;
            this.push(chunk);
            if (i == dataToWrite.length) {
              this.push(null);
            }
          });
          dataToWrite.forEach(data => duplexStream.write(data));
        }
      },
      write(chunk, encoding, callback) {
        // Implement write logic
        events.emit('data', chunk);
        callback();
      },
      final(callback) {
        events.emit('data', null);
        callback();
      },
    });

    const dataToWrite = ['Data1', 'Data2', 'Data3'];
    let receivedData = '';

    duplexStream.on('data', chunk => {
      receivedData += chunk.toString();
    });

    duplexStream.on('end', () => {
      expect(receivedData).to.equal(dataToWrite.join(''));
      done();
    });

    duplexStream.resume();
  });

  it('should handle object mode in both readable and writable sides', done => {
    const events = new EventEmitter();
    let written = false;

    const duplexStream = new Duplex({
      objectMode: true,
      read(size) {
        if (!written) {
          written = true;
          let i = 0;
          events.on('data', (chunk: any) => {
            i++;
            this.push(chunk);
            if (i === dataToWrite.length) {
              this.push(null);
            }
          });
          dataToWrite.forEach(data => duplexStream.write(data));
        }
      },
      write(chunk, encoding, callback) {
        // Implement write logic
        events.emit('data', chunk);
        callback();
      },
      final(callback) {
        events.emit('data', null);
        callback();
      },
    });

    const dataToWrite = [
      { data: 'Data1' },
      { data: 'Data2' },
      { data: 'Data3' },
    ];
    let receivedData: any[] = [];

    duplexStream.on('data', chunk => {
      receivedData.push(chunk);
    });

    duplexStream.on('end', () => {
      try {
        expect(receivedData).to.deep.equal(dataToWrite);
        done();
      } catch (err) {
        done(err);
      }
    });

    duplexStream.resume();
  });

  it('should handle large amounts of data without memory issues', done => {
    const largeDataSize = 1024 * 1024 * 10; // 1 MB
    const largeData = Buffer.alloc(largeDataSize, 'x');

    let i = 0;
    const duplexStream = new Duplex({
      read(size) {
        // Simulate reading data asynchronously
        setTimeout(() => {
          const chunk = largeData.subarray(i, i + size);
          i += size;
          this.push(chunk.length > 0 ? chunk : null);
        }, 0);
      },
      write(chunk, encoding, callback) {
        // Simulate writing data asynchronously
        setTimeout(() => {
          callback();
        }, 10);
      },
    });

    let receivedData = Buffer.alloc(0);

    duplexStream.on('data', chunk => {
      receivedData = Buffer.concat([receivedData, chunk]);
    });

    duplexStream.on('end', () => {
      // Ensure all data is received correctly
      expect(receivedData.length).to.equal(largeDataSize);
      done();
    });

    duplexStream.write('Data1');
    duplexStream.write('Data2');
    duplexStream.resume();
    duplexStream.end();
  }).timeout(5_000); // Adjust the timeout based on the expected completion time

  it('should handle custom events', done => {
    const events = new EventEmitter();
    const duplexStream = new Duplex({
      read(size) {
        // Implement read logic
        events.on('data', this.push);
      },
      write(chunk, encoding, callback) {
        // Implement write logic
        events.emit('data', chunk);
        callback();
      },
    });

    duplexStream.on('customEvent', data => {
      // Handle custom event
      expect(data).to.equal('Custom Data');
      done();
    });

    // Trigger a custom event
    duplexStream.emit('customEvent', 'Custom Data');
  });
});
